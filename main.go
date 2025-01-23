package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/vaalley/codex-backend/api"
	"github.com/vaalley/codex-backend/config"
	"github.com/vaalley/codex-backend/db"
)

func main() {
	startTime := time.Now()

	// Load environment variables
	config.LoadConfig()

	// Connect to MongoDB
	db.ConnectMongo()

	// Get port from environment and if not set, default to 3000
	port := config.GetEnv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize a new Fiber app
	app := fiber.New()

	// Register routes
	api.SetupRoutes(app)

	// Calculate startup time and log it along with the port
	startupTime := time.Since(startTime).Milliseconds()
	log.Printf("🆙 Server starting on port %s (startup took %d ms)", port, startupTime)

	// Start the server
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
