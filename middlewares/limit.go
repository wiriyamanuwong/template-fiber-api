package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	memory "github.com/gofiber/storage/memory/v2"
)

// NewLimit creates a new Limit storage default Memory
func NewLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        100,
		Expiration: 10 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		Storage: memory.New(memory.Config{
			GCInterval: 10 * time.Second,
		}),
	})
}
