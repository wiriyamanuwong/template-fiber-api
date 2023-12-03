package routes

import (
	"strings"

	"github.com/attapon-th/template-fiber-api/controllers/pingctl"
	"github.com/attapon-th/template-fiber-api/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// NewRouters initializes fiber router
func NewRouters(r fiber.Router) {
	prefix := viper.GetString("prefix")
	prefix = strings.TrimSuffix(prefix, "/")

	r.Use("/", middlewares.NewHelmet(),
		middlewares.NewCORS(),
		middlewares.NewLimit(),
		middlewares.NewAccessLog(),
		middlewares.NewCompress(),
	)

	pingctl.NewPingCtl(r.Group(prefix + "/ping"))

	log.Debug().Str("path", prefix+"/public").Msg("Router Public")
	createStaticRoute(r, prefix+"/public")

	log.Debug().Str("path", prefix+"/api").Msg("Router RestAPI")
	createRestAPIRouter(r, prefix+"/api")

	log.Debug().Str("path", prefix+"/swagger").Msg("Router swagger")
	cretateSwagerRoute(r, prefix+"/swagger")
}
