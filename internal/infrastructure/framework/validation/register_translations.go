package validation

import (
	"github.com/gin-gonic/gin/binding"
	enLoc "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
)

// TODO extract this into a factory and let Fx inject it
var (
	Trans ut.Translator
)

func RegisterTranslations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		eng := enLoc.New()
		uni := ut.New(eng, eng)

		var found bool
		Trans, found = uni.GetTranslator("en")
		if !found {
			panic("translator not found")
		}

		if err := en.RegisterDefaultTranslations(v, Trans); err != nil {
			panic(err)
		}

		if err := v.RegisterTranslation("notnull", Trans,
			func(ut ut.Translator) error {
				return ut.Add("notnull", "The {0} field must not be null", true)
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("notnull", fe.Field())
				return t
			},
		); err != nil {
			panic(err)
		}
	}
}
