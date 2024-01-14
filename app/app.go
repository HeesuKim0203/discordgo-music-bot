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
	_, ok := guilds[m.GuildID]

	if !ok {
		guilds[m.GuildID] = util.NewActiveGuild(m.GuildID)
	}

	return guilds[m.GuildID]

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

	inputCommand := content[1]

	switch {
	case strings.EqualFold(inputCommand, commands[Play]):
		if !activeGuild.GetStreamingState() {
			commandHandler.StreamingPlayAndPrepar(s, m, activeGuild)
		} else {
			s.ChannelMessageSend(m.ChannelID, ":exclamation: It's already streaming.")
		}
	case strings.EqualFold(inputCommand, commands[Add]):
		commandHandler.Add(s, m, activeGuild, content[2])
	case strings.EqualFold(inputCommand, commands[View]):
		commandHandler.View(s, m, activeGuild)
	case strings.EqualFold(inputCommand, commands[Delete]):
		commandHandler.Delete(s, m, activeGuild, content[2])
	case strings.EqualFold(inputCommand, commands[Exit]):
		if activeGuild.GetStreamingState() {
			commandHandler.Stop(s, m, activeGuild)
		} else {
			s.ChannelMessageSend(m.ChannelID, ":exclamation: It's not streaming right now!")
		}
	case strings.EqualFold(inputCommand, commands[Skip]):
		if activeGuild.GetStreamingState() {
			commandHandler.Skip(s, m, activeGuild)
		} else {
			s.ChannelMessageSend(m.ChannelID, ":exclamation: It's not streaming right now!")
		}
	case strings.EqualFold(inputCommand, commands[Search]):
		searchText := ""
		for _, v := range content[2:] {
			searchText += v
		}
		commandHandler.Search(s, m, searchText)
	case strings.EqualFold(inputCommand, commands[Help]):
		s.ChannelMessageSend(m.ChannelID, `
__**Commands**__
**search** : Searches for a song. You must enter the song title as text.
**add** : Adds a song to the playlist. You need to enter the song's id. The id of the song can be found in the search results.
**delete** : Deletes a song from the playlist. You must enter the id of the song you want to delete.
**view** : Displays the current playlist and the songs playing in the streaming playlist.
**play** : Streams the current playlist. The streaming playlist cannot be changed.
**exit** : Deletes all songs in the currently streaming playlist.
**skip** : Skips one song.
**help** : Provides a manual for the commands.

If a bug is found, please create an issue at https://github.com/HeesuKim0203/discordgo-music-bot
`)
	default:
		s.ChannelMessageSend(m.ChannelID, ":x: Not Found command!")
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
