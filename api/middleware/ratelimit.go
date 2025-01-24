package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

// RateLimit creates a rate limiter middleware
func RateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        5,               // max count per interval
		Expiration: 1 * time.Minute, // interval for the limit
		KeyGenerator: func(c fiber.Ctx) string {
			return c.IP() // use IP address as key
		},
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, please try again later",
			})
		},
	})
}
