package config

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func NewServer() *fiber.App {
	config := fiber.Config{
		StrictRouting: false,
		CaseSensitive: false,
	}

	app := fiber.New(config)

	app.Use(logger.New())
	ConfigureCors(app)

	return app
}
