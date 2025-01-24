package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

// ValidateRequest validates the request body against the provided struct
func ValidateRequest(payload interface{}) fiber.Handler {
	return func(c fiber.Ctx) error {
		if err := c.Bind().Body(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if err := validate.Struct(payload); err != nil {
			errors := make(map[string]string)
			for _, err := range err.(validator.ValidationErrors) {
				errors[err.Field()] = getErrorMsg(err)
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": errors,
			})
		}

		return c.Next()
	}
}

// getErrorMsg returns a human-readable error message for a given validation error
func getErrorMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	default:
		return "Invalid value"
	}
}
