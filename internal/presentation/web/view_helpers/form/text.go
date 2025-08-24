package form

import (
	"bytes"
	"html/template"
)

type Text struct {
	ID, Label, Name    string
	Placeholder, Value *string
	Required           bool
}

func RenderText(tmpl *template.Template, f Text) (template.HTML, error) {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "#/form/text", f)
	return template.HTML(buf.String()), err
}
