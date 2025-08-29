package mailables

import (
	"app/internal/core/component/user/domain/verification"
)

type PasswordSetupMail interface {
	Render(verification verification.ID, fullName string, token string) (message string, err error)
}
