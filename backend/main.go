package main

import (
	"github.com/woaitsAryan/fampay-task/backend/config"
	"github.com/woaitsAryan/fampay-task/backend/helpers"
	"github.com/woaitsAryan/fampay-task/backend/initializers"
	"github.com/woaitsAryan/fampay-task/backend/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.AddLogger()
	initializers.ConnectToCache()
	initializers.AutoMigrate()

	go initializers.FetchVideos()
}

func main() {
	defer initializers.LoggerCleanUp()

	app := fiber.New(fiber.Config{
		ErrorHandler: helpers.ErrorHandler,
	})

	app.Use(helmet.New())
	app.Use(config.CORS())
	app.Use(logger.New())

	routers.Config(app)

	app.Listen(":" + initializers.CONFIG.PORT)
}