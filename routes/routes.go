package routes

import (
	"codex-backend/controllers"
	"codex-backend/db"
	"codex-backend/middleware"
	"codex-backend/models"
	"codex-backend/repositories"
	"codex-backend/services"

	"github.com/gofiber/fiber/v3"
)

// RegisterRoutes registers API routes
func RegisterRoutes(app *fiber.App) {
	// Initialize dependencies
	platformRepo := repositories.NewMongoPlatformRepository(db.GetCollection("platforms"))
	platformService := services.NewPlatformService(platformRepo)
	platformController := controllers.NewPlatformController(platformService)

	api := app.Group("/api")

	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("It runs!")
	})

	// |-------------------------------------------------------|
	// |                    Platform Routes                    |
	// |-------------------------------------------------------|

	api.Get("/platforms", platformController.GetPlatforms)
	api.Get("/platforms/:id", platformController.GetPlatformByID)
	api.Post("/platforms", middleware.ValidateRequestBody(&models.Platform{}), platformController.CreatePlatform)
	api.Put("/platforms/:id", middleware.ValidateRequestBody(&models.Platform{}), platformController.UpdatePlatform)
	api.Delete("/platforms/:id", platformController.DeletePlatform)
}
