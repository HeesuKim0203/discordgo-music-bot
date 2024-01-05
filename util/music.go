package util

import (
	"fmt"
	"time"
)

type Music struct {
	title     string        // Music title
	id        string        // Music Id
	streamUrl string        // Stream Url
	duration  time.Duration // Music Duration
}

func NewMusic(title string, id string, StreamUrl string, durationSeconds int) *Music {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", durationSeconds))
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
