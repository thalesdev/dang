package config

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func ConfigureCors(app *fiber.App) *fiber.App {
	app.Use(cors.New())
	return app
}
