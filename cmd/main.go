package main

import (
	"codex-backend/config"
	"codex-backend/db"
	"codex-backend/middleware"
	"codex-backend/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Load config and connect to MongoDB
	config.LoadConfig()
	db.ConnectMongo()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandlerMiddleware,
	})

	// Register routes
	routes.RegisterRoutes(app)

	// Get port from environment and if not set, default to 3000
	port := config.GetEnv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server starting on port %s", port)

	// Start server
	log.Fatal(app.Listen(":" + port))
}
