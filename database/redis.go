package database

import (
	"fmt"

	"github.com/gofiber/storage/redis/v3"
)

var store *redis.Storage

func ConnectRedis(config redis.Config) {
	store = redis.New(config)
	fmt.Println("redis connected")
}

func GetRedis() *redis.Storage {
	return store
}
