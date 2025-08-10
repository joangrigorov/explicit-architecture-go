package validation

import (
	"app/internal/infrastructure/framework/validation/rules"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterBindings() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("notnull", rules.NotNull)
		if err != nil {
			panic(err)
		}
	}
}
