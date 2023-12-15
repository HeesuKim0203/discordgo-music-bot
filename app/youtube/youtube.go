package youtube

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var apiKey string

type youtubeServiceHandler struct {
	service *youtube.Service
	options []string
}

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Not Found env : ")
		panic(err)
	}

	apiKey = os.Getenv("YOUTUBE_PUBLIC_KEY")
}

func (y *youtubeServiceHandler) SearchHandle(search string, max int) *youtube.SearchListResponse {

	searchResponse, err := y.service.Search.List(y.options).Q(search).Type("video").MaxResults(int64(max)).Do()

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
