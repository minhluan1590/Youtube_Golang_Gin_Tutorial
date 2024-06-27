package validators

import (
	"strings"
	"github.com/go-playground/validator/v10"
)

func ValidateIsCool(fl validator.FieldLevel) bool {
	// More interesting validation logic: Check if the string contains the word "cool"
	value := fl.Field().String()
	return len(value) > 0 && strings.Contains(strings.ToLower(value), "cool")
}
