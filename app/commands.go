package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-music-bot/util"
	"github.com/discordgo-music-bot/youtube"
)

const (
	Search = "search"
	Play   = "play"
	Stop   = "stop"
	Skip   = "skip"
	Add    = "add"
	View   = "view"
)

type CommandHandler struct {
	youtube  *youtube.YoutubeServiceHandler
	Commands map[string]string // Command text
}

func NewCommandHandler() *CommandHandler {

	commands := make(map[string]string)

	commands[Search] = Search
	commands[Play] = Play
	commands[Stop] = Stop
	commands[Skip] = Skip
	commands[View] = View
	commands[Add] = Add

	commandHandler := &CommandHandler{
		youtube:  youtube.NewService(),
		Commands: commands,
	}

	return commandHandler
}

func (h *CommandHandler) SearchMusic(s *discordgo.Session, m *discordgo.MessageCreate, searchText string) {

	if searchText == "" {
		s.ChannelMessageSend(m.ChannelID, "No text. Please enter text.")
		return
	}

	text := ""
	searchData := h.youtube.SearchHandle(searchText)

	for _, item := range searchData.Items {
		title := item.Snippet.Title
		videoId := item.Id.VideoId

		text += "Title: " + title + "\n"
		text += "Video ID: " + videoId + "\n"
		text += "Video URL: " + "https://www.youtube.com/watch?v=" + videoId + "\n\n"
	}

	s.ChannelMessageSend(m.ChannelID, text)
}

func (h *CommandHandler) AddMusic(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild, musicId string) {

	if musicId == "" {
		s.ChannelMessageSend(m.ChannelID, "No text. Please enter text.")
		return
	}

	text := ""
	data, err := h.youtube.SearchToIdHandle(musicId)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Please enter a valid ID. Search failed.")
		return
	}

	music := util.NewMusic(data[0].Snippet.Title, data[0].Id.VideoId, "https://www.youtube.com/watch?v="+data[0].Id.VideoId)

	ag.EnqueueMedia(music)

	title := music.GetTitle()
	videoId := music.GetId()

	text += "Title: " + title + "\n"
	text += "Video ID: " + videoId + "\n The video has been added."

	s.ChannelMessageSend(m.ChannelID, text)
}

func (h *CommandHandler) StopMusic(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild) {
	ag.GetEvent().Stop()
}

func (h *CommandHandler) SkipMusic(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild) {
	ag.GetEvent().Skip()
}

func (h *CommandHandler) PlayMusic(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild) {

	vc, err := s.ChannelVoiceJoin(m.GuildID, m.ChannelID, false, false)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Voice is not available. Please try on a channel where voice chat is available.")
		return
	}

	defer vc.Disconnect()

	// vc.Speaking(true)
	// defer vc.Speaking(false)

	// client := ytDownload.Client{}
	// video, err := client.GetVideo(youtubeId)
	// if err != nil {
	// 	log.Println("Error getting video info:", err)
	// 	return
	// }

	// // Get Stream
	// formats := video.Formats.WithAudioChannels().FindByQuality("medium")
	// stream, err := client.GetStreamURL(video, formats)
	// if err != nil {
	// 	log.Println("Error getting video stream Url :", err)
	// 	return
	// }

	// // Set options
	// options := dca.StdEncodeOptions
	// options.BufferedFrames = 100
	// options.FrameDuration = 20
	// options.CompressionLevel = 5
	// options.Bitrate = 96

	// // Encode
	// encodeSession, err := dca.EncodeFile(stream, options)
	// if err != nil {
	// 	log.Printf("[%s] Failed to create encoding session for \"%s\": %s", s.State.User.Username, "./video.mp4", err.Error())
	// 	return
	// }
	// defer encodeSession.Cleanup()

	// time.Sleep(500 * time.Millisecond)

	// done := make(chan error)
	// dca.NewStream(encodeSession, vc, done)

	// select {
	// case err := <-done:
	// 	if err != nil && err != io.EOF {
	// 		log.Printf("[%s] Error occurred during stream for \"%s\": %s", s.State.User.Username, "./video.mp4", err.Error())
	// 		return
	// 	}
	// }

	// _ = encodeSession.Stop()
}
