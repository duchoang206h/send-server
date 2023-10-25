package midleware

import (
	"github.com/duchoang206h/send-server/config"
	"github.com/duchoang206h/send-server/database"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Limiter() func(*fiber.Ctx) error { // should use di
	storage := database.GetRedis()
	return limiter.New(config.GetLimiterConfig(storage))
}
