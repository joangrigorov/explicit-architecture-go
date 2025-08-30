package mailables

import (
	"app/internal/core/component/user/application/mailables"
	"bytes"
	"fmt"
	"html/template"
)

type UserConfirmedMail struct {
}

func NewUserConfirmedMail() mailables.UserConfirmedMail {
	return &UserConfirmedMail{}
}

func (c *UserConfirmedMail) Render(fullName string) (message string, err error) {
	tmpl, err := template.
		New(fmt.Sprintf("%T", c)).
		ParseFiles("internal/infrastructure/component/user/mailables/user_confirmed_mail.gohtml")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "mail/user_confirmed", map[string]any{
		"FullName": fullName,
	}); err != nil {
		panic(err)
	}

	return buf.String(), nil
}
