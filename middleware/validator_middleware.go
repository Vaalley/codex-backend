package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// ErrorHandlerMiddleware to handle validation errors
func ErrorHandlerMiddleware(c fiber.Ctx, err error) error {
	// Check if it's a validation error
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, e.Error())
		}
		return c.Status(400).JSON(fiber.Map{"errors": errorMessages})
	}

	// Fallback to the default error handler
	return c.Status(500).SendString("Internal Server Error")
}
