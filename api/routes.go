package api

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"

	"github.com/vaalley/codex-backend/api/middleware"
	"github.com/vaalley/codex-backend/handlers"
	"github.com/vaalley/codex-backend/models"
)

// SetupRoutes registers the routes of the app
func SetupRoutes(app *fiber.App) {
	log.Println("🔄 Registering routes... 🛣️")

	//  ---------------
	// | Public routes |
	//  ---------------
	// "/livez" Checks if the server is up and running
	app.Get(healthcheck.DefaultLivenessEndpoint, healthcheck.NewHealthChecker())
	// "/readyz" Assesses if the application is ready to handle requests
	app.Get(healthcheck.DefaultReadinessEndpoint, healthcheck.NewHealthChecker())
	// "/startupz" Checks if the application has completed its startup sequence and is ready to proceed with initialization and readiness checks
	app.Get(healthcheck.DefaultStartupEndpoint, healthcheck.NewHealthChecker())
	// app.Get("/metrics", monitor.New()) // Doesn't work because github.com/gofiber/fiber/v3/middleware/monitor doesn't exist?

	// Auth routes with rate limiting and validation
	auth := app.Group("/auth", middleware.RateLimit())
	auth.Post("/login", middleware.ValidateRequest(&models.LoginRequest{}), handlers.Login)
	auth.Post("/register", middleware.ValidateRequest(&models.RegisterRequest{}), handlers.Register)
	auth.Post("/logout", middleware.JWTAuth(), handlers.Logout)

	//  ------------------
	// | Protected routes |
	//  ------------------
	api := app.Group("/api", middleware.APIKeyAuth())

	api.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// Games routes
	//     games := api.Group("/games")
	//     games.Get("/", handlers.GetGames)
	//     games.Post("/", handlers.CreateGame)
	//     games.Get("/:id", handlers.GetGame)
	//     games.Put("/:id", handlers.UpdateGame)
	//     games.Delete("/:id", handlers.DeleteGame)

	log.Println("✅ Routes registered successfully")
}
