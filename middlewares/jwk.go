package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

const LocalsUserToken = "user-token"

// NewJWK creates a new JWK middleware
func NewJWK() fiber.Handler {
	return jwtware.New(jwtware.Config{
		JWKSetURLs:     []string{viper.GetString("JWK_URL")},
		SuccessHandler: nil,
		ContextKey:     LocalsUserToken,
	})
}
