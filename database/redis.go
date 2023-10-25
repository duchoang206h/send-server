package database

import (
	"fmt"

	"github.com/duchoang206h/send-server/config"
	"github.com/gofiber/storage/redis/v3"
)

var store *redis.Storage

func ConnectRedis() {
	store = redis.New(config.GetRedisConfig())
	fmt.Println("redis connected")
}

func GetRedis() *redis.Storage {
	return store
}
