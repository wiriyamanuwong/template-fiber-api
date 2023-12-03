package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func NewHelmet() fiber.Handler {
	return helmet.New()
}
