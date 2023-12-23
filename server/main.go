package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
    logger "github.com/JyJy32/logger"
)

type VideoInfo struct {
    videoId   string
    title     string
    channel   string
    thumbnail string
}

func main() {
    logger.Debug("Starting server")
    loadenv()
    if os.Getenv("API_KEY") == "" {
        panic("API_KEY is not set")
    }

    server_init()
}

func loadenv() {
   if err := godotenv.Load(); err != nil {
      log.Println("No .env file found")
   }
}

func get_video(videoID string) (*youtube.Video, error) {
    apikey := os.Getenv("API_KEY")
    ctx := context.Background();

    service, err := youtube.NewService(ctx, option.WithAPIKey(apikey));
    if err != nil {
        log.Fatalf("Error creating new YouTube client: %v", err)
        return nil, err
    }

    response, err := service.Videos.List([]string{"id", "snippet", "contentDetails", "statistics"}).Id(videoID).Do();
    if err != nil {
        log.Fatalf("Error calling Videos.List: %v", err)
        return nil, err
    }
    return response.Items[0], nil 
} 

