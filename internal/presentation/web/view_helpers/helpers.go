package view_helpers

import (
	"app/internal/presentation/web/services/session"
	"app/internal/presentation/web/view_helpers/form"
	"html/template"
	"log"
)

func Helpers(tmpl *template.Template) template.FuncMap {
	return template.FuncMap{
		"gui_form_text": func(required bool, name string, label string, values map[string]interface{}, placeholder string, errors map[string]string) template.HTML {
			html, err := form.RenderText(tmpl, form.Text{
				ID:          name,
				Label:       label,
				Name:        name,
				Placeholder: &placeholder,
				Value:       value(values, name),
				Required:    required,
				Error:       errors[name],
			})
			if err != nil {
				return "<!-- error rendering text field -->"
			}
			return html
		},
		"gui_form_password": func(required bool, name string, label string, placeholder string, errors map[string]string) template.HTML {
			log.Println("Password errors", errors)
			html, err := form.RenderPassword(tmpl, form.Password{
				ID:          name,
				Label:       label,
				Name:        name,
				Placeholder: &placeholder,
				Required:    required,
				Error:       errors[name],
			})
			if err != nil {
				return "<!-- error rendering password field -->"
			}
			return html
		},
		"gui_form_email": func(required bool, name string, label string, values map[string]interface{}, placeholder string, errors map[string]string) template.HTML {
			html, err := form.RenderEmail(tmpl, form.Email{
				ID:          name,
				Label:       label,
				Name:        name,
				Placeholder: &placeholder,
				Value:       value(values, name),
				Required:    required,
				Error:       errors[name],
			})
			if err != nil {
				return "<!-- error rendering email field -->"
			}
			return html
		},
		"gui_form_hidden": func(required bool, name string, values map[string]interface{}) template.HTML {
			html, err := form.RenderHidden(tmpl, form.Hidden{
				ID:       name,
				Name:     name,
				Value:    value(values, name),
				Required: required,
			})
			if err != nil {
				return "<!-- error rendering hidden field -->"
			}
			return html
		},
		"gui_alerts": func(alerts []session.Alert) template.HTML {
			html, err := form.RenderAlerts(tmpl, form.Alerts{
				Alerts: alerts,
			})
			if err != nil {
				return "<!-- error rendering alerts -->"
			}
			return html
		},
	}
}

func value(values map[string]interface{}, name string) *string {
	value := values[name]
	if value != nil {
		v := value.(string)
		return &v
	}
	log.Println("No value for "+name+" found", values)
	return nil
}
