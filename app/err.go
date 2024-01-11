package app

import "github.com/bwmarrin/discordgo"

func VoiceJoinErr(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, ":x: Voice is not available. Please try on a channel where voice chat is available.")
	return
}

func NoTextErr(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, ":x: Not Found text. Please enter text.")
	return
}

func SearchFailErr(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, ":x: Search failed.")
	return
}
