package web

import (
	"app/internal/presentation/web/view_helpers"
	"html/template"
)

func NewTemplate() *template.Template {
	tmpl := template.New("")
	tmpl.Funcs(view_helpers.Helpers(tmpl))

	template.Must(tmpl.ParseGlob("internal/presentation/web/view_helpers/**/*.gohtml"))
	template.Must(tmpl.ParseGlob("internal/presentation/web/layout/*.gohtml"))
	template.Must(tmpl.ParseGlob("internal/presentation/web/pages/**/**/*.gohtml"))

	return tmpl
}
