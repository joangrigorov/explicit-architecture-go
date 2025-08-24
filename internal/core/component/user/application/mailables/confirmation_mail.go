package mailables

import (
	"app/internal/core/component/user/domain"
)

type ConfirmationMail interface {
	Render(confirmationID domain.ConfirmationID, fullName string, hmacSum string) (message string, err error)
}
