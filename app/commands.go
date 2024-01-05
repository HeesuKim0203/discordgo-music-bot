package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-music-bot/youtube"
)

type Command struct {
	id      string                                             // Command text
	handler func(*discordgo.Session, *discordgo.MessageCreate) // Command handler
}

type CommandHandler struct {
	youtube       *youtube.YoutubeServiceHandler
	searchText    string
	addMusicId    string
	deleteMusicId string
	commands      *map[string]Command
}

func (h *CommandHandler) setSearchText(searchText string) {
	h.searchText = searchText
}

func (h *CommandHandler) setAddMusicId(musicId string) {
	h.addMusicId = musicId
}

func (h *CommandHandler) setDeleteMusicId(musicId string) {
	h.deleteMusicId = musicId
}

func (h *CommandHandler) searchMusic(s *discordgo.Session, m *discordgo.MessageCreate) {

	if h.searchText == "" {
		s.ChannelMessageSend(m.ChannelID, "No text. Please enter text.")
		return
	}

	text := ""
	searchData := h.youtube.SearchHandle(h.searchText)

	for _, item := range searchData.Items {
		title := item.Snippet.Title
		videoId := item.Id.VideoId

		text += "Title: " + title + "\n"
		text += "Video ID: " + videoId + "\n"
		text += "Video URL: " + "https://www.youtube.com/watch?v=" + videoId + "\n\n"
	}

	s.ChannelMessageSend(m.ChannelID, text)
	h.searchText = ""

}

func (h *CommandHandler) addMusic(s *discordgo.Session, m *discordgo.MessageCreate) {

	if h.addMusicId == "" {
		s.ChannelMessageSend(m.ChannelID, "No text. Please enter text.")
		return
	}

	text := ""
	searchData := h.youtube.SearchToIdHandle(h.addMusicId)

	// Todo : Not found search data!

	s.ChannelMessageSend(m.ChannelID, text)
	h.addMusicId = ""
}

func (h *CommandHandler) deleteMusic(s *discordgo.Session, m *discordgo.MessageCreate) {

	if h.deleteMusicId == "" {
		s.ChannelMessageSend(m.ChannelID, "No text. Please enter text.")
		return
	}

	text := ""
	searchData := h.youtube.SearchToIdHandle(h.deleteMusicId)

	// Todo : Not found search data!

	s.ChannelMessageSend(m.ChannelID, text)
	h.deleteMusicId = ""
}

func (h *CommandHandler) Play(s *discordgo.Session, m *discordgo.MessageCreate) {
	// vc, _ := s.ChannelVoiceJoin(m.GuildID, m.ChannelID, false, false)
	// defer vc.Disconnect()

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
