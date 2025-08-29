package mailables

import (
	"app/internal/core/component/user/domain/confirmation"
)

type PasswordSetupMail interface {
	Render(confirmationID confirmation.ID, fullName string, hmacSum string) (message string, err error)
}
