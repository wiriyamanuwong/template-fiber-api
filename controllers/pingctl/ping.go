package pingctl

import "github.com/gofiber/fiber/v2"

// NewPingCtl creates a new ping controller
func NewPingCtl(r fiber.Router) {
	r.Get("/", ping)
	r.Post("/", ping)
	r.Head("/", ping)
	// r.All("/", ping)
}

func ping(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
