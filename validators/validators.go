package validators

import (
	"strings"

	"github.com/go-playground/validator"
)

func ValidateContains(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "B")
}
