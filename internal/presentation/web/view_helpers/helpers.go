package view_helpers

import (
	"app/internal/presentation/web/view_helpers/form"
	"html/template"
)

func Helpers(tmpl *template.Template) template.FuncMap {
	return template.FuncMap{
		"gui_form_text": func(required bool, name string, label string, value string, placeholder string) template.HTML {
			html, err := form.RenderText(tmpl, form.Text{
				ID:          name,
				Label:       label,
				Name:        name,
				Placeholder: &placeholder,
				Value:       &value,
				Required:    required,
			})
			if err != nil {
				return "<!-- error rendering text field -->"
			}
			return html
		},
		"gui_form_password": func(required bool, name string, label string, placeholder string) template.HTML {
			html, err := form.RenderPassword(tmpl, form.Password{
				ID:          name,
				Label:       label,
				Name:        name,
				Placeholder: &placeholder,
				Required:    required,
			})
			if err != nil {
				return "<!-- error rendering password field -->"
			}
			return html
		},
		"gui_form_email": func(required bool, name string, label string, value string, placeholder string) template.HTML {
			html, err := form.RenderEmail(tmpl, form.Email{
				ID:          name,
				Label:       label,
				Name:        name,
				Placeholder: &placeholder,
				Value:       &value,
				Required:    required,
			})
			if err != nil {
				return "<!-- error rendering password field -->"
			}
			return html
		},
	}
}
