package mailables

import (
	"app/config/api"
	"app/internal/core/component/user/application/mailables"
	"app/internal/core/component/user/domain"
	"bytes"
	"fmt"
	"html/template"
)

const templateFile = "internal/infrastructure/component/user/mailables/confirmation_mail.gohtml"

type ConfirmationMail struct {
	webURL string
}

func NewConfirmationMail(cfg *api.Config) mailables.ConfirmationMail {
	return &ConfirmationMail{webURL: cfg.Frontend.WebURL}
}

func (c *ConfirmationMail) Render(
	confirmationID domain.ConfirmationID,
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
	if err := tmpl.ExecuteTemplate(&buf, "mail/confirmation", map[string]interface{}{
		"WebURL":         c.webURL,
		"FullName":       fullName,
		"HmacSum":        hmacSum,
		"ConfirmationID": confirmationID.String(),
	}); err != nil {
		panic(err)
	}

	return buf.String(), nil
}
