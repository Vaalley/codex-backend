package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vaalley/codex-backend/utils"
)

// JWTAuth returns a middleware that validates JWT tokens
func JWTAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Try to get token from cookie first
		tokenString := c.Cookies("session_token")

		// Fallback to Authorization header
		if tokenString == "" {
			authHeader := c.Get("Authorization")
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				tokenString = authHeader[7:]
			}
		}

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authentication token",
			})
		}

		// Validate JWT
		token, err := utils.ParseJWT(tokenString)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		// Store user information in context
		c.Locals("userID", claims["sub"])
		c.Locals("userRoles", claims["roles"])

		return c.Next()
	}
}
