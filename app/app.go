package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-music-bot/dca"
	"github.com/discordgo-music-bot/youtube"
	"github.com/joho/godotenv"

	you "github.com/kkdai/youtube/v2"
)

var y = youtube.NewService()

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Split(m.Content, " ")
	search := ""

	if content[0] != "!letsgobot" {
		return
	}

	if content[1] == "play" {

		vc, _ := s.ChannelVoiceJoin(m.GuildID, m.ChannelID, false, false)

		vc.Speaking(true)

		search = content[2]
		client := you.Client{}

		video, err := client.GetVideo(search)
		if err != nil {
			log.Println("Error getting video info:", err)
			return
		}
		defer vc.Disconnect()

		formats := video.Formats.WithAudioChannels().FindByQuality("medium")
		stream, _, err := client.GetStream(video, formats)
		if err != nil {
			log.Println("Error getting video stream:", err)
			return
		}

		file, err := os.Create("./video.mp4")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.Copy(file, stream)
		if err != nil {
			panic(err)
		}

		options := dca.StdEncodeOptions
		options.BufferedFrames = 100
		options.FrameDuration = 20
		options.CompressionLevel = 5
		options.Bitrate = 96

		encodeSession, err := dca.EncodeFile("./video.mp4", options)
		if err != nil {
			log.Printf("[%s] Failed to create encoding session for \"%s\": %s", s.State.User.Username, "./video.mp4", err.Error())
			return
		}
		defer encodeSession.Cleanup()

		time.Sleep(500 * time.Millisecond)

		done := make(chan error)
		dca.NewStream(encodeSession, vc, done)

		select {
		case err := <-done:
			if err != nil && err != io.EOF {
				log.Printf("[%s] Error occurred during stream for \"%s\": %s", s.State.User.Username, "./video.mp4", err.Error())
				return
			}
		}

		_ = encodeSession.Stop()
		return
	}

	if content[1] != "" {

		text := ""
		searchData := y.SearchHandle(content[2], 10)

		for _, item := range searchData.Items {
			title := item.Snippet.Title
			videoId := item.Id.VideoId

			text += "Title: " + title + "\n"
			text += "Video ID: " + videoId + "\n"
			text += "Video URL: " + "https://www.youtube.com/watch?v=" + videoId + "\n\n"
		}

		s.ChannelMessageSend(m.ChannelID, text)
	}
}

func NewDiscord() *discordgo.Session {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Not Found env : ")
		panic(err)
	}

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_PUBLIC_KEY"))

	if err != nil {
		fmt.Println("discord Create Error : ")
		panic(err)
	}

	// Logging Server
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	discord.AddHandler(messageCreate)

	return discord
}
