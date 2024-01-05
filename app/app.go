package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-music-bot/config"
)

var (
	c       = config.GetConfig()
	botName = c.GetBotName()
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Split(m.Content, " ")

	if content[0] != botName {
		return
	}

	// switch content[1] {
	// case "play":
	// 	Play(s, m, content[2])
	// 	return
	// case "add":
	// 	return
	// case "search":
	// 	searchText := ""
	// 	for _, v := range content[2:] {
	// 		searchText += v
	// 	}
	// 	Search(s, m, searchText)
	// 	return
	// default:
	// 	s.ChannelMessageSend(m.ChannelID, "Not Found Command!")
	// 	return
	// }
}

func NewDiscord() *discordgo.Session {

	discord, err := discordgo.New("Bot " + c.GetDiscordToken())

	if err != nil {
		fmt.Println("discord Create Error : ")
		panic(err)
	}

	// Logging Server
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	discord.AddHandler(messageHandler)

	return discord
}
