package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/woaitsAryan/fampay-task/backend/controllers"
)

func VideosRouter(app *fiber.App) {
	app.Get("/videos", controllers.GetVideos)
}
