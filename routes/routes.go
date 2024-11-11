package routes

import (
	"codex-backend/controllers"

	"github.com/gofiber/fiber/v3"
)

// RegisterRoutes registers API routes
func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// |-------------------------------------------------------|
	// |                    Platform Routes                    |
	// |-------------------------------------------------------|

	// Get a list of all platforms
	api.Get("/get-platforms", controllers.GetPlatforms)
	// Get a specific platform by its ID (sent as a query parameter)
	api.Get("/get-platform-by-id", controllers.GetPlatformByID)

	// api.Post("/create-platform", controllers.CreatePlatform)
}
