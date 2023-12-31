package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-music-bot/config"
	"github.com/discordgo-music-bot/util"
)

var (
	c              = config.GetConfig()
	botName        = c.GetBotName()
	guilds         = make(map[string]*util.ActiveGuild)
	commandHandler = NewCommandHandler()
	commands       = commandHandler.Commands
)

func checkGuild(s *discordgo.Session, m *discordgo.MessageCreate) *util.ActiveGuild {

	// Verify that a guild exists
	activeGuild, ok := guilds[m.GuildID]

	if !ok {
		activeGuild = util.NewActiveGuild(m.GuildID)
	}

	return activeGuild

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	activeGuild := checkGuild(s, m)

	content := strings.Split(m.Content, " ")

	if !strings.EqualFold(content[0], botName) {
		return
	}

	switch {
	case strings.EqualFold(content[1], commands[Play]):
		// Play(s, m, content[2])
	case strings.EqualFold(content[1], commands[Add]):
		commandHandler.Add(s, m, activeGuild, content[2])
	case strings.EqualFold(content[1], commands[View]):
		commandHandler.View(s, m, activeGuild)
	case strings.EqualFold(content[1], commands[Delete]):
	case strings.EqualFold(content[1], commands[Stop]):
		commandHandler.Stop(s, m, activeGuild)
	case strings.EqualFold(content[1], commands[Skip]):
		commandHandler.Skip(s, m, activeGuild)
	case strings.EqualFold(content[1], commands[Search]):
		searchText := ""
		for _, v := range content[2:] {
			searchText += v
		}
		commandHandler.Search(s, m, searchText)
	default:
		s.ChannelMessageSend(m.ChannelID, "Not Found Command!")
	}

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
