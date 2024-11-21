package middleware

import (
	"codex-backend/config"

	"github.com/gofiber/fiber/v3"
)

// ValidateAPIKey middleware checks if the request has a valid API key
func ValidateAPIKey(c fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")
	expectedAPIKey := config.GetEnv("API_KEY")

	if apiKey == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "API key is missing",
		})
	}

	if apiKey != expectedAPIKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid API key",
		})
	}

	return c.Next()
}
