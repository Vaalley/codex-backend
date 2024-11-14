package utils

import (
	"github.com/go-playground/validator/v10"
)

// Initialize the validator
var validate = validator.New()

// ValidateStruct validates a struct and returns an error if validation fails
func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}
