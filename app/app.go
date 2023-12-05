package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "")
	flag.Parse()
}

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

	discord, err := discordgo.New("Bot " + os.Getenv("PUBLIC_KEY"))

	if err != nil {
		fmt.Println("discord Create Error : ")
		panic(err)
	}

	discord.AddHandler(messageCreate)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	return discord
}
