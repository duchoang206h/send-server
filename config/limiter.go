package config

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/redis/v3"
)

func GetLimiterConfig(storage *redis.Storage) limiter.Config {
	limiter_max, _ := strconv.Atoi(Config("LIMITER_MAX"))
	limiter_expiration, _ := strconv.Atoi(Config("LIMITER_EXPIRATION_SECOND"))
	config := limiter.Config{
		Max:               limiter_max,
		Expiration:        time.Duration(limiter_expiration) * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"message": "Too Many Requests"})
		},
		Storage: storage,
	}
	return config
}
