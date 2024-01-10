package app

import "github.com/bwmarrin/discordgo"

func voiceJoinErr(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, ":x:Voice is not available. Please try on a channel where voice chat is available.")
	return
}
