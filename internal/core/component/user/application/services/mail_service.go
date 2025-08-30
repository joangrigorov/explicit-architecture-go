package services

import (
	"app/internal/core/component/user/application/mailables"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/component/user/domain/verification"
	"app/internal/core/port/errors"
	"app/internal/core/port/mail"
)

type MailService struct {
	mailer            mail.Mailer
	passwordSetupMail mailables.PasswordSetupMail
	userConfirmedMail mailables.UserConfirmedMail
	errors            errors.ErrorFactory
}

func NewMailService(
	mailer mail.Mailer,
	passwordSetupMail mailables.PasswordSetupMail,
	userConfirmedMail mailables.UserConfirmedMail,
	errors errors.ErrorFactory,
) *MailService {
	return &MailService{
		mailer:            mailer,
		passwordSetupMail: passwordSetupMail,
		userConfirmedMail: userConfirmedMail,
		errors:            errors,
	}
}

func (s *MailService) SendUserConfirmedMail(
	recipientEmail user.Email,
	senderEmail user.Email,
	fullName string,
) error {
	message, err := s.userConfirmedMail.Render(fullName)
	if err != nil {
		return s.errors.New(errors.ErrMail, "Error rendering user confirmed mail", err)
	}

	to := []string{recipientEmail.String()}
	var cc []string

	const subject = "You completed your Activity Planner account!"

	if err := s.mailer.Send(to, cc, cc, senderEmail.String(), subject, message); err != nil {
		return s.errors.New(errors.ErrMail, "Error sending user confirmed mail", err)
	}
	
	return nil
}

func (s *MailService) SendPasswordSetupMail(
	verificationID verification.ID,
	recipientEmail user.Email,
	senderEmail user.Email,
	fullName string,
	token string,
) error {
	message, err := s.passwordSetupMail.Render(verificationID, fullName, token)

	if err != nil {
		return s.errors.New(errors.ErrMail, "Error rendering password setup mail", err)
	}

	to := []string{recipientEmail.String()}
	var cc []string

	const subject = "Verify your Activity Planner account"

	if err := s.mailer.Send(to, cc, cc, senderEmail.String(), subject, message); err != nil {
		return s.errors.New(errors.ErrMail, "Error sending password setup mail", err)
	}

	return nil
}
