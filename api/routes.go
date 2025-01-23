package api

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"

	"github.com/vaalley/codex-backend/api/middleware"
)

// Registers the routes of the app
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
	//     app.Post("/login", handlers.Login) // need to add handlers.Login
	//     app.Post("/register", handlers.Register) // need to add handlers.Register

	//  ------------------
	// | Protected routes |
	//  ------------------
	api := app.Group("/api", middleware.Auth())

	// Test route
	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
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
