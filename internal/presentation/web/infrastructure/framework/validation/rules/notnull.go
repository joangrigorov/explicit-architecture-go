package rules

import (
	"github.com/aarondl/null/v9"
	. "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type NotNull struct{}

func (v *NotNull) Tag() string {
	return "notnull"
}

func (v *NotNull) Validate(fl validator.FieldLevel) bool {
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

func (v *NotNull) RegisterTranslations(tr Translator) error {
	return tr.Add("notnull", "The {0} field must not be null", true)
}

func (v *NotNull) Translate(tr Translator, fe validator.FieldError) string {
	t, _ := tr.T("notnull", fe.Field())
	return t
}

func NewNotNull() *NotNull {
	return &NotNull{}
}
