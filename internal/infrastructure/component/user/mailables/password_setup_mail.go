package mailables

import (
	"app/config/api"
	"app/internal/core/component/user/application/mailables"
	"app/internal/core/component/user/domain/confirmation"
	"bytes"
	"fmt"
	"html/template"
)

const templateFile = "internal/infrastructure/component/user/mailables/password_setup_mail.gohtml"

type PasswordSetupMail struct {
	webURL string
}

func NewPasswordSetupMail(cfg *api.Config) mailables.PasswordSetupMail {
	return &PasswordSetupMail{webURL: cfg.Frontend.WebURL}
}

func (c *PasswordSetupMail) Render(
	confirmationID confirmation.ID,
	fullName string,
	hmacSum string,
) (message string, err error) {
	tmpl, err := template.
		New(fmt.Sprintf("%T", c)).
		ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "mail/password_setup", map[string]interface{}{
		"WebURL":         c.webURL,
		"FullName":       fullName,
		"HmacSum":        hmacSum,
		"ConfirmationID": confirmationID.String(),
	}); err != nil {
		panic(err)
	}

	return buf.String(), nil
}
