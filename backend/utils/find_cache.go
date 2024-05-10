package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/woaitsAryan/fampay-task/backend/cache"
	"github.com/woaitsAryan/fampay-task/backend/models"
	schemas "github.com/woaitsAryan/fampay-task/backend/schema"
)

func FindCache(urlParams schemas.VideoFetchSchema, ctx context.Context) ([]models.Video, error) {

	cacheKey := fmt.Sprintf("%s-%d-%d", urlParams.Title, urlParams.Limit, urlParams.Page)

	cachedResult, err := cache.GetFromCache(cacheKey, ctx)
	if err == nil {
		return cachedResult, nil
	}
	return nil, errors.New("cache miss")
}