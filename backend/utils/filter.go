package utils

import (
	schemas "github.com/woaitsAryan/fampay-task/backend/schema"
	"gorm.io/gorm"
)

func FilterVideos(db *gorm.DB, urlParams schemas.VideoFetchSchema) *gorm.DB {

	query := db

	query = query.Limit(urlParams.Limit)

	if urlParams.Title != "" {
		query = db.Where("title LIKE ?", "%"+urlParams.Title+"%")
	}

	query = query.Offset((urlParams.Page - 1) * 20)
	
	return query
}
