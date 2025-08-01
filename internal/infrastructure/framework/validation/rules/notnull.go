package rules

import (
	"github.com/aarondl/null/v9"
	"github.com/go-playground/validator/v10"
)

func NotNull(fl validator.FieldLevel) bool {
	n, ok := fl.Field().Interface().(null.Value)
	if !ok {
		return false // wrong type
	}

	if !n.IsSet() {
		return true
	}

	if n.IsValid() {
		return true
	}

	return false
}
