package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/woaitsAryan/fampay-task/backend/helpers"
	"github.com/woaitsAryan/fampay-task/backend/initializers"
	"github.com/woaitsAryan/fampay-task/backend/models"
)

func GetFromCache(key string, ctx context.Context) ([]models.Video, error) {
	data, err := initializers.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("item not found in cache")
		}
		go helpers.LogServerError("Error Getting from cache", err, "")
		return nil, fmt.Errorf("error getting from cache")
	}
	var videos []models.Video
	err = json.Unmarshal([]byte(data), &videos)
	if err != nil {
		go helpers.LogServerError("Error unmarshalling data from cache", err, "")
		return nil, fmt.Errorf("error unmarshalling data from cache")
	}

	return videos, nil
}

func SetToCache(key string, data []models.Video, ctx context.Context) error {

	jsonData, err := json.Marshal(data)
    if err != nil {
        go helpers.LogServerError("Error marshalling data to JSON", err, "")
        return fmt.Errorf("error marshalling data to JSON")
    }

	if err := initializers.RedisClient.Set(ctx, key, jsonData, initializers.CacheExpirationTime).Err(); err != nil {
		go helpers.LogServerError("Error Setting to cache", err, "")
		return fmt.Errorf("error setting to cache")
	}
	return nil
}
