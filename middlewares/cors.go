package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// NewCORS creates a new CORS
func NewCORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "http://localhost",
	})
}
