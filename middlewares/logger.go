package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// NewAccessLog creates a new AccessLog
func NewAccessLog() fiber.Handler {
	return logger.New(logger.Config{
		Format:        "[${time}][${method}][${status}][${latency}][${ip}] - ${url}\n",
		TimeFormat:    "2006-01-02T15:04:05.999Z07:00",
		TimeZone:      "Asia/Bangkok",
		DisableColors: false,
		Output:        os.Stdout,
	})
}
