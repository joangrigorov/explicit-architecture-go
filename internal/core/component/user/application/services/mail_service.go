package services

import (
	"app/internal/core/component/user/application/mailables"
	"app/internal/core/component/user/domain"
	"app/internal/core/port/mail"
)

type MailService struct {
	mailer      mail.Mailer
	confirmMail mailables.ConfirmationMail
}

func NewMailService(mailer mail.Mailer, confirmMail mailables.ConfirmationMail) *MailService {
	return &MailService{mailer: mailer, confirmMail: confirmMail}
}

func (s *MailService) SendConfirmationMail(
	confirmationID domain.ConfirmationID,
	recipientEmail string,
	senderEmail string,
	fullName string,
	hmac string,
) error {
	message, err := s.confirmMail.Render(confirmationID, fullName, hmac)

	if err != nil {
		return err
	}

	strings := []string{recipientEmail}
	var cc []string

	const subject = "Confirm your Activity Planner account"

	if err := s.mailer.Send(strings, cc, cc, senderEmail, subject, message); err != nil {
		return err
	}

	return nil
}
