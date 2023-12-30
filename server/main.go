package main

import (
	"context"
	"log"
	"os"

	logger "github.com/JyJy32/logger"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

type VideoInfo struct {
	VideoId   string
	Title     string
	Channel   string
	Thumbnail string
}

var currentVideo VideoInfo

func main() {
	logger.Debug("Starting server")
	loadenv()
	if os.Getenv("API_KEY") == "" {
		panic("API_KEY is not set. Please set the API_KEY environment variable or in a .env file")
	}

	server_init()
}

func loadenv() {
	if err := godotenv.Load(); err != nil {
		logger.Debug("No .env file found")
	}
}

func getVideoInfo(videoID string) (*VideoInfo, error) {
	apikey := os.Getenv("API_KEY")
	ctx := context.Background()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apikey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
		return nil, err
	}

	response, err := service.Videos.List([]string{"id", "snippet"}).Id(videoID).Do()
	if err != nil {
		log.Fatalf("Error calling Videos.List: %v", err)
		return nil, err
	}
	videoInfo := &VideoInfo{
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Snippet.ChannelTitle,
		response.Items[0].Snippet.Thumbnails.Default.Url,
	}
	return videoInfo, nil
}
