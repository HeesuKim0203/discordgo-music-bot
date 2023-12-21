package app

import (
	"io"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-music-bot/dca"
	you "github.com/kkdai/youtube/v2"
)

func Play(s *discordgo.Session, m *discordgo.MessageCreate, youtubeId string) {
	vc, _ := s.ChannelVoiceJoin(m.GuildID, m.ChannelID, false, false)

	vc.Speaking(true)
	defer vc.Speaking(false)

	client := you.Client{}
	video, err := client.GetVideo(youtubeId)
	if err != nil {
		log.Println("Error getting video info:", err)
		return
	}
	defer vc.Disconnect()

	// Get Stream
	formats := video.Formats.WithAudioChannels().FindByQuality("medium")
	stream, err := client.GetStreamURL(video, formats)
	if err != nil {
		log.Println("Error getting video stream Url :", err)
		return
	}

	// Set options
	options := dca.StdEncodeOptions
	options.BufferedFrames = 100
	options.FrameDuration = 20
	options.CompressionLevel = 5
	options.Bitrate = 96

	// Encode
	encodeSession, err := dca.EncodeFile(stream, options)
	if err != nil {
		log.Printf("[%s] Failed to create encoding session for \"%s\": %s", s.State.User.Username, "./video.mp4", err.Error())
		return
	}
	defer encodeSession.Cleanup()

	time.Sleep(500 * time.Millisecond)

	done := make(chan error)
	dca.NewStream(encodeSession, vc, done)

	select {
	case err := <-done:
		if err != nil && err != io.EOF {
			log.Printf("[%s] Error occurred during stream for \"%s\": %s", s.State.User.Username, "./video.mp4", err.Error())
			return
		}
	}

	_ = encodeSession.Stop()
}

func Search(s *discordgo.Session, m *discordgo.MessageCreate, youtubeSearchText string) {

	text := ""
	searchData := y.SearchHandle(youtubeSearchText, 5)

	for _, item := range searchData.Items {
		title := item.Snippet.Title
		videoId := item.Id.VideoId

		text += "Title: " + title + "\n"
		text += "Video ID: " + videoId + "\n"
		text += "Video URL: " + "https://www.youtube.com/watch?v=" + videoId + "\n\n"
	}

	s.ChannelMessageSend(m.ChannelID, text)

}
