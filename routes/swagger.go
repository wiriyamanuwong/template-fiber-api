package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func cretateSwagerRoute(r fiber.Router, prerfix string) {

	swaggerConfig := swagger.ConfigDefault
	r.Get(prerfix+"/*", swagger.New(swaggerConfig)) // default
}
