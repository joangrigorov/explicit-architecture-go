package web

import (
	"app/internal/presentation/web/view_helpers"
	"html/template"

	ut "github.com/go-playground/universal-translator"
)

func NewTemplate(tr ut.Translator) *template.Template {
	tmpl := template.New("")
	tmpl.Funcs(view_helpers.Helpers(tmpl, tr))

	template.Must(tmpl.ParseGlob("internal/presentation/web/view_helpers/**/*.gohtml"))
	template.Must(tmpl.ParseGlob("internal/presentation/web/layout/*.gohtml"))
	template.Must(tmpl.ParseGlob("internal/presentation/web/pages/**/**/*.gohtml"))

	return tmpl
}
