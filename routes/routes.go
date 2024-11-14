package routes

import (
	"codex-backend/controllers"

	"github.com/gofiber/fiber/v3"
)

// RegisterRoutes registers API routes
func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("It runs!")
	})

	// |-------------------------------------------------------|
	// |                    Platform Routes                    |
	// |-------------------------------------------------------|

	api.Get("/platforms", controllers.GetPlatforms)
	api.Get("/platforms/:id", controllers.GetPlatformByID)
	api.Post("/platforms", controllers.CreatePlatform)
}
