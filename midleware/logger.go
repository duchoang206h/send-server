package midleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() func(*fiber.Ctx) error {
	return logger.New()
}
