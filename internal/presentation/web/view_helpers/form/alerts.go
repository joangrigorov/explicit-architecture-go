package form

import (
	"app/internal/presentation/web/services/session"
	"bytes"
	"html/template"
)

type Alerts struct {
	Alerts []session.Alert
}

func RenderAlerts(tmpl *template.Template, a Alerts) (template.HTML, error) {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "#/form/alerts", a)
	return template.HTML(buf.String()), err
}
