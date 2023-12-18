package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/lets-go-bot/youtube"

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

		search = content[2]
		client := you.Client{}

		log.Println(search)

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

		cmd := exec.Command("ffmpeg", "-i", "./video.mp4", "-f", "s16le", "-ar", "48000", "-ac", "2", "-")

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Println("Error creating audio pipe:", err)
			return
		}

		err = cmd.Start()
		if err != nil {
			log.Println("Error starting ffmpeg:", err)
			return
		}

		vc.Speaking(true)
		defer vc.Speaking(false)

		buf := make([]byte, 2048)
		for {
			select {
			case <-time.After(250 * time.Millisecond):
				_, err := stdout.Read(buf)
				if err != nil {
					log.Println("Error reading audio data:", err)
					return
				}
				vc.OpusSend <- buf
			}
		}

		//buf := make([]byte, 2048)

		// for {
		// 	select {
		// 	case <-time.After(250 * time.Millisecond):
		// 		_, err := stream.Read(buf)
		// 		log.Println(buf)
		// 		if err != nil {
		// 			log.Println("Error reading audio data:", err)
		// 			return
		// 		}
		// 		vc.OpusSend <- buf
		// 	}
		// }
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
