package util

import (
	"github.com/discordgo-music-bot/config"
)

type ActiveGuild struct {
	id          string      // Guild Id
	music       []*Music    // Music Queue
	musicChan   chan *Music // Music Chan
	event       *Event      // Guild Event Management
	isStreaming bool        // State Streaming
}

func NewActiveGuild(id string) *ActiveGuild {

	config := config.GetConfig()

	maxSize := config.GetQueueSize()

	newEvnet := NewEvent()

	return &ActiveGuild{
		id:          id,
		music:       make([]*Music, maxSize),
		musicChan:   make(chan *Music, maxSize),
		isStreaming: false,
		event:       newEvnet,
	}
}

func (ag *ActiveGuild) MusicIsFulling() bool {
	return len(ag.musicChan) == cap(ag.musicChan)
}

func (ag *ActiveGuild) CheckExistenceMusicQueue() bool {
	return !(len(ag.musicChan) == 0)
}

func (ag *ActiveGuild) GetMusicChanSize() int {
	return len(ag.musicChan)
}

func (ag *ActiveGuild) GetStreamingState(isStreaming bool) bool {
	return ag.isStreaming
}

func (ag *ActiveGuild) GetMusicChan() chan *Music {
	return ag.musicChan
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

func (ag *ActiveGuild) PreparStreaming() {
	for _, item := range ag.music {
		ag.musicChan <- item
	}
}

func (ag *ActiveGuild) CleanUp() {
	close(ag.musicChan)
	ag.isStreaming = false
	ag.event = nil
	ag.musicChan = nil
}
