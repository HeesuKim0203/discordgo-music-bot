package util

import (
	"fmt"
	"time"

	"github.com/discordgo-music-bot/config"
)

type Music struct {
	title     string        // Music title
	id        string        // Music Id
	streamUrl string        // Stream Url
	duration  time.Duration // Music Duration
}

var c = config.GetConfig()

func NewMusic(title string, id string, StreamUrl string) *Music {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", c.GetMusicDuration()))
	return &Music{
		title:     title,
		id:        id,
		streamUrl: StreamUrl,
		duration:  duration,
	}
}

func (m *Music) GetTitle() string {
	return m.title
}

func (m *Music) GetId() string {
	return m.id
}

func (m *Music) GetStreamUrl() string {
	return m.streamUrl
}

func (m *Music) GetDuration() time.Duration {
	return m.duration
}
