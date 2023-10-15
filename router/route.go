package router

import (
	"github.com/duchoang206h/send-server/handler"
	"github.com/duchoang206h/send-server/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	// inject dependencies
	fileRepo:= repository.NewFileRepository()
	fileHandler:= handler.NewFileHandler(fileRepo)

	api:= app.Group("/api")
	fileRoute := api.Group("/file")
	//todo: rate limit
	fileRoute.Post("/", fileHandler.HandleFileUpload)
	fileRoute.Get("/:hash", fileHandler.HandleDownloadFile)

}
