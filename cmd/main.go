package main

import (
	"log"

	"github.com/duchoang206h/send-server/config"
	"github.com/duchoang206h/send-server/database"
	"github.com/duchoang206h/send-server/middleware"
	"github.com/duchoang206h/send-server/router"
	"github.com/gofiber/fiber/v2"
)

// ReadCloserWrapper is a custom implementation of io.ReadCloser
func main() {
	app := fiber.New()
	mongoURI := config.GetMongoURI()
	dbName := config.GetMongoDBName()
	err := database.ConnectMongo(mongoURI, dbName)
	// redisConfig := config.GetRedisConfig()
	// database.ConnectRedis(redisConfig)
	if err != nil {
		log.Fatal(err)
	}
	app.Use(middleware.Logger())
	// app.Use(middleware.Limiter())
	router.SetupRoute(app)
	log.Fatal(app.Listen(config.Config("SERVER_PORT")))
}
