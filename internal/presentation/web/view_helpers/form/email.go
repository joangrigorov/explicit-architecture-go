package form

import (
	"bytes"
	"html/template"
)

type Email struct {
	ID, Label, Name    string
	Placeholder, Value *string
	Required           bool
}

func RenderEmail(tmpl *template.Template, f Email) (template.HTML, error) {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "#/form/email", f)
	return template.HTML(buf.String()), err
}
