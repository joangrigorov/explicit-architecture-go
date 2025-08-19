package validation

import (
	"app/internal/infrastructure/validation/rules"

	"github.com/gin-gonic/gin/binding"
	enLoc "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
)

// Rule represents a validation rule including the validator function and translations
type Rule interface {

	// Tag is unique key for using the validation rule
	Tag() string

	// Validate is the validation function itself
	Validate(fl validator.FieldLevel) bool

	// RegisterTranslations adds translations for the validator
	RegisterTranslations(tr ut.Translator) error

	// Translate is the translation function itself
	Translate(tr ut.Translator, fe validator.FieldError) string
}

// List of custom validation rules
var customRules = []Rule{
	rules.NewNotNull(),
}

// RegisterRules bootstraps the custom validation rules
func RegisterRules(translator ut.Translator, v *validator.Validate) {
	for _, rule := range customRules {
		tag := rule.Tag()

		if err := v.RegisterValidation(tag, rule.Validate); err != nil {
			panic(err)
		}
		if err := en.RegisterDefaultTranslations(v, translator); err != nil { // coverage-ignore
			panic(err)
		}

		if err := v.RegisterTranslation(
			tag,
			translator,
			rule.RegisterTranslations,
			rule.Translate,
		); err != nil { // coverage-ignore
			panic(err)
		}
	}
}

// NewTranslator provides the ut.Translator factory
func NewTranslator() ut.Translator {
	locale := enLoc.New()
	uni := ut.New(locale, locale)

	t, _ := uni.GetTranslator(locale.Locale())
	return t
}

// NewValidatorValidate registers the validator.Validate
// type so it can be used in dependency injection
func NewValidatorValidate() *validator.Validate {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		return v
	}
	panic("Cannot register validator")
}
