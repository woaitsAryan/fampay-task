package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/woaitsAryan/fampay-task/backend/cache"
	"github.com/woaitsAryan/fampay-task/backend/helpers"
	"github.com/woaitsAryan/fampay-task/backend/initializers"
	"github.com/woaitsAryan/fampay-task/backend/models"
	"github.com/woaitsAryan/fampay-task/backend/utils"
)

func GetVideos(c *fiber.Ctx) error {

	urlParams, err := helpers.ValidateVideoURLParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Query Params",
		})
	}

	cachedVideos, err := utils.FindCache(*urlParams, c.Context())
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"videos": cachedVideos,
		})
	}

	db := initializers.DB

	db = utils.FilterVideos(db, *urlParams)

	var videos []models.Video

	if err := db.Find(&videos).Error; err != nil {
		helpers.LogDatabaseError("Error getting videos", err, "controllers/get_videos.go")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error getting videos",
		})
	}
	
	cacheKey := fmt.Sprintf("%s-%d-%d", urlParams.Title, urlParams.Limit, urlParams.Page)

	cache.SetToCache(cacheKey, videos, c.Context())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"videos": videos,
	})
}
