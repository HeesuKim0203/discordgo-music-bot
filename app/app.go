package app

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/lets-go-bot/app/youtube"
)

var y = youtube.NewService()

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Split(m.Content, " ")

	if content[0] != "!letsgobot" {
		return
	}

	if content[1] != "" {

		text := ""
		searchData := y.SearchHandle(content[1], 15)

		for _, item := range searchData.Items {
			title := item.Snippet.Title
			videoId := item.Id.VideoId

			text += "Title: " + title + ", Video ID: " + videoId + "\n"
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
