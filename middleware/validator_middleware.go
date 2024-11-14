package middleware

import (
	"codex-backend/models"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// getErrorMsg returns a user-friendly error message for validation errors
func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		if fe.Type().Kind() == reflect.String {
			return fmt.Sprintf("Must be at least %s characters long", fe.Param())
		}
		return fmt.Sprintf("Must be at least %s", fe.Param())
	case "max":
		if fe.Type().Kind() == reflect.String {
			return fmt.Sprintf("Must not exceed %s characters", fe.Param())
		}
		return fmt.Sprintf("Must not exceed %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", fe.Param())
	}
	return "Invalid value"
}

// ErrorHandlerMiddleware to handle validation errors
func ErrorHandlerMiddleware(c fiber.Ctx, err error) error {
	var errors []models.ValidationError

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			field := strings.ToLower(err.Field())
			errors = append(errors, models.ValidationError{
				Field:   field,
				Message: getErrorMsg(err),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Validation failed",
			Errors:  errors,
		})
	}

	// Handle other types of errors
	return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
		Status:  fiber.StatusInternalServerError,
		Message: "Internal Server Error",
	})
}

// ValidateRequestBody middleware to validate request body
func ValidateRequestBody(schema interface{}) fiber.Handler {
	return func(c fiber.Ctx) error {
		if err := json.Unmarshal(c.Body(), schema); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Status:  fiber.StatusBadRequest,
				Message: "Invalid request body",
			})
		}

		if err := validate.Struct(schema); err != nil {
			return ErrorHandlerMiddleware(c, err)
		}

		return c.Next()
	}
}
