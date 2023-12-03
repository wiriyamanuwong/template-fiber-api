// Package pkg is the vaccine-hpv-api package
package pkg

import (
	"errors"
	"runtime"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// FiberConfig is the Fiber configuration Version 2
type FiberConfig struct {
	app *fiber.App
	fiber.Config
	Port int
	Host string
}

func jsonMarshal(v interface{}) ([]byte, error) {
	return json.MarshalWithOption(v, json.UnorderedMap())
}

// NewFiber creates a new Fiber
func NewFiber(config ...fiber.Config) *FiberConfig {
	var cfg fiber.Config
	if len(config) == 0 {
		cfg = fiber.Config{
			AppName:                      "API Service",
			ServerHeader:                 "API Service",
			BodyLimit:                    1 * 1024 * 1024, // 1MB
			ReadBufferSize:               10 * 1024,
			ReadTimeout:                  10 * time.Second,
			WriteTimeout:                 30 * time.Second,
			Prefork:                      false,
			DisablePreParseMultipartForm: true,
			DisableStartupMessage:        true,
			CaseSensitive:                true,
			EnablePrintRoutes:            true,
			Views:                        nil,
			JSONEncoder:                  jsonMarshal,
			JSONDecoder:                  json.Unmarshal,
			ErrorHandler:                 FiberErrorHandler,
		}
	} else {
		cfg = config[0]
	}

	viper.SetDefault("prefork", 1)
	viper.SetDefault("port", 8888)
	viper.SetDefault("host", "0.0.0.0")
	if n := viper.GetInt("prefork"); n > 1 {
		cfg.Prefork = true
		maxProcessing := runtime.GOMAXPROCS(n)
		if maxProcessing > n {
			maxProcessing = n
		}
		log.Info().Bool("prefork", cfg.Prefork).Int("maxProcessing", maxProcessing).Msg("Prefork enabled")
	}

	return &FiberConfig{
		app:    fiber.New(cfg),
		Port:   viper.GetInt("port"),
		Host:   viper.GetString("host"),
		Config: cfg,
	}
}

// Fiber returns the FiberConfig
func (c *FiberConfig) Fiber() *fiber.App {
	return c.app
}

// Listen starts the Fiber
func (c *FiberConfig) Listen() error {
	log.Info().Str("host", c.Host).Int("port", c.Port).Msg("Server started")
	return c.app.Listen(c.Host + ":" + strconv.Itoa(c.Port))
}

// FiberErrorHandler handles errors
func FiberErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		msg = e.Error()
	}
	c.Set("Content-Type", fiber.MIMEApplicationJSONCharsetUTF8)

	// Send custom error page
	_ = c.Status(code).JSON(fiber.Map{
		"code":    code,
		"ok":      false,
		"message": msg,
	})

	// Return from handler
	return nil
}
