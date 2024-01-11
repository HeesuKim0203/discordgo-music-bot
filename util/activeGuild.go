package util

import (
	"github.com/discordgo-music-bot/config"
)

type ActiveGuild struct {
	id          string   // Guild Id
	music       []*Music // Music Current PlayList
	playMusic   []*Music // Music Playing PlayList
	event       *Event   // Guild Event Management
	isStreaming bool     // State Streaming
	maxSize     int
}

func NewActiveGuild(id string) *ActiveGuild {

	config := config.GetConfig()

	maxSize := config.GetQueueSize()

	event := NewEvent()

	return &ActiveGuild{
		id:          id,
		isStreaming: false,
		event:       event,
		maxSize:     maxSize,
	}
}

func (ag *ActiveGuild) MusicIsFulling() bool {
	return len(ag.music) >= ag.maxSize
}

func (ag *ActiveGuild) GetStreamingState() bool {
	return ag.isStreaming
}

func (ag *ActiveGuild) GetPlayMusic() []*Music {
	return ag.playMusic
}

func (ag *ActiveGuild) GetEvent() *Event {
	return ag.event
}

func (ag *ActiveGuild) SetStreamingState(isStreaming bool) bool {
	ag.isStreaming = isStreaming
	return ag.isStreaming
}

func (ag *ActiveGuild) GetMusic() []*Music {
	return ag.music
}

func (ag *ActiveGuild) AddMusic(music *Music) {
	ag.music = append(ag.music, music)
}

func (ag *ActiveGuild) DeleteMusic(num int) []*Music {
	ag.music = append(ag.music[:num], ag.music[num+1:]...)

	return ag.music
}

func (ag *ActiveGuild) PreparStreaming() []*Music {
	ag.playMusic = append(ag.playMusic, ag.music...)

	return ag.playMusic
}

func (ag *ActiveGuild) CleanUp() {
	var musicArr []*Music

	ag.isStreaming = false
	ag.event = NewEvent()
	ag.playMusic = musicArr
}
