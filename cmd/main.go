package main

import (
	"log"

	"github.com/duchoang206h/send-server/config"
	"github.com/duchoang206h/send-server/database"
	"github.com/duchoang206h/send-server/router"
	"github.com/gofiber/fiber/v2"
)

// ReadCloserWrapper is a custom implementation of io.ReadCloser
func main () {
	app:=fiber.New()
	err:= database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	router.SetupRoute(app)
	log.Fatal(app.Listen(config.Config("SERVER_PORT")))
}
