package config

import (
	"runtime"
	"strconv"

	"github.com/gofiber/storage/redis/v3"
)

func GetRedisConfig() redis.Config {
	host := Config("REDIS_HOST")
	port, _ := strconv.Atoi(Config("REDIS_HOST"))
	username := Config("REDIS_USERNAME")
	password := Config("REDIS_PASSWORD")
	database, _ := strconv.Atoi(Config("REDIS_DATABASE"))

	config := redis.Config{
		Host:      host,
		Port:      port,
		Username:  username,
		Password:  password,
		Database:  database,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	}
	return config
}
