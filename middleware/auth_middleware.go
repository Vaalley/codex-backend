package middleware

import (
	"codex-backend/config"
	"log"

	"github.com/gofiber/fiber/v3"
)

// ValidateAPIKey middleware checks if the request has a valid API key
func ValidateAPIKey(c fiber.Ctx) error {
	path := c.Path()
	ip := c.IP()
	apiKey := c.Get("X-API-Key")
	expectedAPIKey := config.GetEnv("API_KEY")

	if apiKey == "" {
		log.Printf(" Auth Failed: Missing API key - IP: %s, Path: %s", ip, path)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "API key is missing",
		})
	}

	if apiKey != expectedAPIKey {
		log.Printf(" Auth Failed: Invalid API key - IP: %s, Path: %s", ip, path)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid API key",
		})
	}

	log.Printf(" Auth Success - IP: %s, Path: %s", ip, path)
	return c.Next()
}
