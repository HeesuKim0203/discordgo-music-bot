package youtube

import (
	"context"
	"log"

	"github.com/discordgo-music-bot/config"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	c             = config.GetConfig()
	apiKey        = c.GetYoutubeToken()
	searchListMax = c.GetMusicSearchListMax()
)

type youtubeServiceHandler struct {
	service *youtube.Service
	options []string
}

func (y *youtubeServiceHandler) SearchHandle(search string) *youtube.SearchListResponse {

	searchResponse, err := y.service.Search.List(y.options).Q(search).Type("video").MaxResults(int64(searchListMax)).Do()

	if err != nil {
		panic(err)
	}

	return searchResponse
}

func (y *youtubeServiceHandler) SearchToIdHandle(searchId string) *youtube.SearchListResponse {

	searchResponse, err := y.service.Search.List(y.options).Q(searchId).Type("video").MaxResults(1).Do()

	if err != nil {
		panic(err)
	}

	return searchResponse

}

func (y *youtubeServiceHandler) SetOption(option string) {
	y.options = append(y.options, option)
}

func NewService() *youtubeServiceHandler {

	ctx := context.Background()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	y := &youtubeServiceHandler{service: service}

	y.SetOption("snippet")

	return y
}
