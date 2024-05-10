package initializers

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/woaitsAryan/fampay-task/backend/config"
	"github.com/woaitsAryan/fampay-task/backend/models"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchVideos() {
	ctx := context.Background()
	for {
		fetchAndInsertVideos(ctx)
		time.Sleep(config.DelayTime)
	}
}

func fetchAndInsertVideos(ctx context.Context) {
	service, err := youtube.NewService(ctx, option.WithAPIKey(CONFIG.DeveloperKeys[0]))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	call := service.Search.List([]string{"id", "snippet"}).
		Q(config.Query).
		MaxResults(config.MaxResults).Type("video").Order("date").PublishedAfter(time.Now().Add(-10 * time.Minute).Format(time.RFC3339))

	response, err := call.Do()
	if err != nil {
		if strings.Contains(err.Error(), "quotaExceeded"){
			CONFIG.RotateDeveloperKey()
			return
		} else{
			log.Fatal(err.Error())
		}
		log.Fatal(err.Error())
	}

	tx := DB.Begin()

	for _, item := range response.Items {
		video := models.Video{
			Title:        item.Snippet.Title,
			VideoID:      item.Id.VideoId,
			Description:  item.Snippet.Description,
			PublishedAt:  item.Snippet.PublishedAt,
			ThumbnailURL: item.Snippet.Thumbnails.High.Url,
		}
		if err := tx.Create(&video).Error; err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	tx.Commit()

	log.Println("Fetched Videos")
}
