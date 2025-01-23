package middleware

import (
	"crypto/sha256"
	"crypto/subtle"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
	"github.com/vaalley/codex-backend/config"
)

// Validates the API key
func validateAPIKey(c fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(config.GetEnv("API_KEY")))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

// Returns a middleware that validates API keys
func Auth() fiber.Handler {
	return keyauth.New(keyauth.Config{
		Validator: validateAPIKey,
		KeyLookup: "header:X-API-Key",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or missing API key",
			})
		},
	})
}
