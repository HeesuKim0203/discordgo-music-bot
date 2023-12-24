package youtube

import (
	"log"
	"testing"
)

func TestYoutubeServiceSearch(t *testing.T) {
	y := NewService()

	searchResponse := y.SearchHandle("lin nax x")

	for _, item := range searchResponse.Items {
		title := item.Snippet.Title
		videoId := item.Id.VideoId
		log.Printf("Title: %s, Video ID: %s\n", title, videoId)
	}
}
