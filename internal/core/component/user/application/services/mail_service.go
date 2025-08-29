package services

import (
	"app/internal/core/component/user/application/mailables"
	"app/internal/core/component/user/domain/confirmation"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/port/mail"
)

type MailService struct {
	mailer            mail.Mailer
	passwordSetupMail mailables.PasswordSetupMail
}

func NewMailService(mailer mail.Mailer, passwordSetupMail mailables.PasswordSetupMail) *MailService {
	return &MailService{mailer: mailer, passwordSetupMail: passwordSetupMail}
}

func (s *MailService) SendPasswordSetupMail(
	confirmationID confirmation.ID,
	recipientEmail user.Email,
	senderEmail user.Email,
	fullName string,
	hmac string,
) error {
	message, err := s.passwordSetupMail.Render(confirmationID, fullName, hmac)

	if err != nil {
		return err
	}

	to := []string{recipientEmail.String()}
	var cc []string

	const subject = "Confirm your Activity Planner account"

	if err := s.mailer.Send(to, cc, cc, senderEmail.String(), subject, message); err != nil {
		return err
	}

	return nil
}
