package main

import (
	"log"

	"github.com/duchoang206h/send-server/config"
	"github.com/duchoang206h/send-server/database"
	"github.com/duchoang206h/send-server/midleware"
	"github.com/duchoang206h/send-server/router"
	"github.com/gofiber/fiber/v2"
)

// ReadCloserWrapper is a custom implementation of io.ReadCloser
func main() {
	app := fiber.New()
	err := database.ConnectMongo()
	database.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}
	app.Use(midleware.Logger())
	app.Use(midleware.Limiter())
	router.SetupRoute(app)
	log.Fatal(app.Listen(config.Config("SERVER_PORT")))
}
