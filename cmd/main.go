package main

import (
	"log"

	"github.com/duchoang206h/send-server/config"
	"github.com/duchoang206h/send-server/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupRoute(app)
	log.Fatal(app.Listen(config.Config("APP_PORT")))
}
