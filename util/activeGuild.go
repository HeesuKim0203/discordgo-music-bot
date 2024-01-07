package util

import (
	"github.com/discordgo-music-bot/config"
)

// Todo : Queue Channel -> Array
type ActiveGuild struct {
	id          string      // Guild Id
	musicQue    chan *Music // Music Queue
	event       *Event      // Guild Event Management
	isStreaming bool        // State Streaming
}

func NewActiveGuild(id string) *ActiveGuild {

	config := config.GetConfig()

	maxSize := config.GetQueueSize()

	newEvnet := NewEvent()

	return &ActiveGuild{
		id:          id,
		musicQue:    make(chan *Music, maxSize),
		isStreaming: false,
		event:       newEvnet,
	}
}

func (ag *ActiveGuild) MusicQueueIsFulling() bool {
	return len(ag.musicQue) == cap(ag.musicQue)
}

func (ag *ActiveGuild) CleanUp() {
	close(ag.musicQue)
	ag.isStreaming = false
	ag.event = nil
	ag.musicQue = nil
}

func (ag *ActiveGuild) CheckExistenceMusicQueue() bool {
	return !(len(ag.musicQue) == 0)
}

func (ag *ActiveGuild) GetMusicQueueSize() int {
	return len(ag.musicQue)
}

func (ag *ActiveGuild) GetStreamingState(isStreaming bool) bool {
	return ag.isStreaming
}

func (ag *ActiveGuild) GetMusicQueue() chan *Music {
	return ag.musicQue
}

func (ag *ActiveGuild) GetEvent() *Event {
	return ag.event
}

func (ag *ActiveGuild) SetStreamingState(isStreaming bool) bool {
	ag.isStreaming = isStreaming
	return ag.isStreaming
}

func (g *ActiveGuild) EnqueueMedia(music *Music) {
	g.musicQue <- music
}
