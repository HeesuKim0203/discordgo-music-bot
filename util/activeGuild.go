package util

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ActiveGuild struct {
	id          string      // Guild Id
	musicQue    chan *Music // Music Queue
	event       *Event      // Guild Event Management
	isStreaming bool        // State Streaming
}

func (ag *ActiveGuild) NewActiveGuild(id string) *ActiveGuild {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Not Found env : ")
		panic(err)
	}

	maxSize, err := strconv.Atoi(os.Getenv("MUSIC_QUEUE_SIZE"))

	if err != nil {
		fmt.Println("Failed convert string -> number : ")
		panic(err)
	}

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
