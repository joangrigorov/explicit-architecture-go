package form

import (
	"bytes"
	"html/template"
)

type Hidden struct {
	ID, Name string
	Value    *string
	Required bool
}

func RenderHidden(tmpl *template.Template, f Hidden) (template.HTML, error) {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "#/form/hidden", f)
	return template.HTML(buf.String()), err
}
