// Package routes for project api router
package routes

import (
	"net/url"

	"github.com/attapon-th/template-fiber-api/controllers/todoctl"
	"github.com/attapon-th/template-fiber-api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Init initializes fiber router
func createRestAPIRouter(r fiber.Router, prefix string) {

	r.Route(prefix+"/v1", func(r fiber.Router) {
		todoctl.NewTodoCtl(r.Group("/todos")) // prefix+/api/v1/todos
	})
	configAPISwagger(prefix + "/v1")
}

func configAPISwagger(apiPrefix string) {
	if baseURL, err := url.Parse(viper.GetString("base_url")); err == nil {
		// Custom Swagger URL
		docs.SwaggerInfo.BasePath = apiPrefix
		docs.SwaggerInfo.Host = baseURL.Host
		docs.SwaggerInfo.Schemes = []string{baseURL.Scheme}
		log.Debug().Str("host", baseURL.Host).Str("schema", baseURL.Scheme).Str("prefix", apiPrefix).Msg("set swagger url.")
	}
}
