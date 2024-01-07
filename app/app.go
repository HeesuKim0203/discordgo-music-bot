package app

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-music-bot/config"
	"github.com/discordgo-music-bot/util"
)

var (
	c              = config.GetConfig()
	botName        = c.GetBotName()
	guilds         = make(map[string]*util.ActiveGuild)
	guildsMutex    = sync.RWMutex{}
	commandHandler = NewCommandHandler()
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

	//activeGuild := checkGuild(s, m)

	content := strings.Split(m.Content, " ")

	if content[0] != botName {
		return
	}

	switch content[1] {
	case commandHandler.Commands[Play]:
		// Play(s, m, content[2])
		return
	case commandHandler.Commands[Add]:
		return
	case commandHandler.Commands[Delete]:
		return
	case commandHandler.Commands[View]:
		return
	case commandHandler.Commands[Stop]:
		return
	case commandHandler.Commands[Skip]:
		return
	case commandHandler.Commands[Search]:
		searchText := ""
		for _, v := range content[2:] {
			searchText += v
		}
		commandHandler.SearchMusic(s, m, searchText)
		return
	default:
		s.ChannelMessageSend(m.ChannelID, "Not Found Command!")
		return
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
