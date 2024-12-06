package main

import (
	"codex-backend/config"
	"codex-backend/db"
	"codex-backend/middleware"
	"codex-backend/routes"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	startTime := time.Now()

	// Load config and connect to MongoDB
	log.Println("⚙️  Loading configuration...")
	config.LoadConfig()

	log.Println("🔌 Connecting to MongoDB...")
	db.ConnectMongo()

	// Initialize Fiber app
	log.Println("🚀 Initializing Fiber application...")
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandlerMiddleware,
	})

	// Register routes
	log.Println("🛣️  Registering routes...")
	routes.RegisterRoutes(app)

	// Get port from environment and if not set, default to 3000
	port := config.GetEnv("PORT")
	if port == "" {
		port = "3000"
	}
	
	startupTime := time.Since(startTime).Seconds()
	log.Printf("✨ Server starting on port %s (startup took %.2f seconds)", port, startupTime)

	// Start server
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
