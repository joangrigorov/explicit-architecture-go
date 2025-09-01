package form

import (
	"bytes"
	"html/template"
)

type Password struct {
	ID, Label, Name string
	Placeholder     *string
	Required        bool
	Error           string
}

func RenderPassword(tmpl *template.Template, f Password) (template.HTML, error) {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "#/form/password", f)
	return template.HTML(buf.String()), err
}
