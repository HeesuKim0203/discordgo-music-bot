package app

import (
	"io"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-music-bot/dca"
	"github.com/discordgo-music-bot/util"
	"github.com/discordgo-music-bot/youtube"
	ytDownload "github.com/kkdai/youtube/v2"
)

const (
	Search = "search"
	Play   = "play"
	Stop   = "stop"
	Skip   = "skip"
	Delete = "delete"
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
	commands[Delete] = Delete
	commands[View] = View
	commands[Add] = Add

	commandHandler := &CommandHandler{
		youtube:  youtube.NewService(),
		Commands: commands,
	}

	return commandHandler
}

func (h *CommandHandler) Search(s *discordgo.Session, m *discordgo.MessageCreate, searchText string) {

	if searchText == "" {
		NoTextErr(s, m)
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

func (h *CommandHandler) Add(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild, musicId string) {

	if musicId == "" {
		NoTextErr(s, m)
		return
	}

	text := ""
	data, err := h.youtube.SearchToIdOrURLHandle(musicId)

	if err != nil {
		SearchFailErr(s, m)
		return
	}

	music := util.NewMusic(data[0].Snippet.Title, data[0].Id.VideoId, "https://www.youtube.com/watch?v="+data[0].Id.VideoId)

	ag.AddMusic(music)

	title := music.GetTitle()
	videoId := music.GetId()

	text += ":green_circle: The video has been added.\n\n"
	text += "Title: " + title + "\n"
	text += "Video ID: " + videoId + "\n"
	text += "Video URL: " + "https://www.youtube.com/watch?v=" + videoId + "\n"

	s.ChannelMessageSend(m.ChannelID, text)
}

func (h *CommandHandler) View(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild) {

	text := ""
	music := ag.GetMusic()
	text += "\n:memo: Current Playlist\n\n"

	text = checkMusicList(music, text)

	playMusic := ag.GetPlayMusic()
	text += "\n:musical_notes: Playlist being played\n\n"

	text = checkMusicList(playMusic, text)

	s.ChannelMessageSend(m.ChannelID, text)
}

func (h *CommandHandler) Delete(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild, musicId string) {

	text := ""
	music := ag.GetMusic()
	findIndex := -1

	for index, item := range music {
		if item.GetId() == musicId {
			findIndex = index
		}
	}

	if findIndex == -1 {
		SearchFailErr(s, m)
		return
	}

	text += ":green_circle: Update Complete.\n\n"

	music = ag.DeleteMusic(findIndex)
	text += ":memo: Current Playlist\n\n"

	text = checkMusicList(music, text)

	s.ChannelMessageSend(m.ChannelID, text)

}

func (h *CommandHandler) Stop(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild) {
	ag.GetEvent().Stop()
}

func (h *CommandHandler) Skip(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild) {
	ag.GetEvent().Skip()
}

func (h *CommandHandler) StreamingPlayAndPrepar(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild) {

	log.Println("Play")

	vc, err := s.ChannelVoiceJoin(m.GuildID, m.ChannelID, false, true)
	if err != nil {
		VoiceJoinErr(s, m)
		return
	}
	defer vc.Disconnect()

	err = vc.Speaking(true)
	if err != nil {
		VoiceJoinErr(s, m)
		return
	}
	defer vc.Speaking(false)

	ag.SetStreamingState(true)
	defer func() {
		ag.CleanUp()
	}()

	playMusic := ag.PreparStreaming()
	for _, item := range playMusic {

		log.Println("item", item)

		if !vc.Ready {
			VoiceJoinErr(s, m)
			return
		}

		if !ag.GetEvent().GetStopState() {
			h.play(s, m, ag, vc, item)
		}

		// Music Delay
		time.Sleep(time.Second)
	}

	log.Println("End")
}

func (h *CommandHandler) play(s *discordgo.Session, m *discordgo.MessageCreate, ag *util.ActiveGuild, vc *discordgo.VoiceConnection, music *util.Music) {

	log.Println(music.GetStreamUrl())

	client := ytDownload.Client{}
	video, err := client.GetVideo(music.GetStreamUrl())

	if err != nil {
		log.Println("Error getting video info:", err)
		return
	}

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
		log.Println("Encode Error")
		return
	}
	defer encodeSession.Cleanup()

	// Encode Session Delay
	time.Sleep(500 * time.Millisecond)

	done := make(chan error)
	dca.NewStream(encodeSession, vc, done)
	defer encodeSession.Stop()

	// Todo : Session Code Debug
	select {
	case err := <-done:
		if err != nil && err != io.EOF {
			log.Println("EOF")
			return
		}
	case <-ag.GetEvent().GetSkipEvent():
		log.Println("GetSkipEvent")
		return
	case <-ag.GetEvent().GetStopEvent():
		log.Println("GetStopEvent")
		return
	}
}

func checkMusicList(music []*util.Music, text string) string {

	if len(music) == 0 {
		text += ":x: Music does not exist.\n"
	} else {
		for _, item := range music {
			title := item.GetTitle()
			videoId := item.GetId()

			text += "Title: " + title + "\n"
			text += "Video ID: " + videoId + "\n"
		}
	}

	return text
}
