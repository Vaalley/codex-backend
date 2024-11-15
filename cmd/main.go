package main

import (
	"codex-backend/config"
	"codex-backend/db"
	"codex-backend/middleware"
	"codex-backend/routes"
	"codex-backend/utils"
	"log"
	"strconv"

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

	// Find available port starting from 3000
	port := utils.FindAvailablePort(3000)
	log.Printf("Server starting on port %d", port)

	// Start server
	log.Fatal(app.Listen(":" + strconv.Itoa(port)))
}
