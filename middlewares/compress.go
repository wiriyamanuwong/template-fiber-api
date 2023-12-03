package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

// NewCompress creates a new Compress middleware. use Accept-Encoding: gzip, ...
func NewCompress() fiber.Handler {
	return compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	})
}
