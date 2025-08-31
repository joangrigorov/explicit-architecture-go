package form

import (
	"bytes"
	"errors"
	"html/template"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Error struct {
	TargetField string
	FormErrors  error
}

func RenderError(tmpl *template.Template, tr ut.Translator, f Error) (template.HTML, error) {
	var ve validator.ValidationErrors
	if errors.As(f.FormErrors, &ve) {
		for _, e := range ve {
			if e.Field() == f.TargetField {
				var buf bytes.Buffer
				err := tmpl.ExecuteTemplate(&buf, "#/form/error", map[string]interface{}{
					"Error": e.Translate(tr),
				})
				return template.HTML(buf.String()), err
			}
		}
	}

	return "", nil
}
