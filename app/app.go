package app

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func NewDiscord() *discordgo.Session {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Not Found env :")
		panic(err)
	}

	discord, err := discordgo.New("Bot " + os.Getenv("PUBLIC_KEY"))

	if err != nil {
		fmt.Println("discord Create Error : ")
		panic(err)
	}

	discord.AddHandler(messageCreate)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	return discord
}
