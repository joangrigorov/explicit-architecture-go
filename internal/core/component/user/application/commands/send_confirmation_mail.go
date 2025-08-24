package commands

import (
	"app/internal/core/component/user/application/errors"
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/application/services"
	sk "app/internal/core/shared_kernel/domain"
	"context"
	"encoding/json"
)

type SendConfirmationMailCommand struct {
	userID      string
	senderEmail string
}

func NewSendConfirmationMailCommand(userID string, senderEmail string) SendConfirmationMailCommand {
	return SendConfirmationMailCommand{userID: userID, senderEmail: senderEmail}
}

func (s SendConfirmationMailCommand) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userID":      s.userID,
		"senderEmail": s.senderEmail,
	})
}

type SendConfirmationMailHandler struct {
	userRepository      repositories.UserRepository
	confirmationService *services.ConfirmationService
	mailService         *services.MailService
}

func NewSendConfirmationMailHandler(
	userRepository repositories.UserRepository,
	confirmationService *services.ConfirmationService,
	mailService *services.MailService,
) *SendConfirmationMailHandler {
	return &SendConfirmationMailHandler{
		userRepository:      userRepository,
		confirmationService: confirmationService,
		mailService:         mailService,
	}
}

func (s *SendConfirmationMailHandler) Handle(ctx context.Context, c SendConfirmationMailCommand) error {
	userID := sk.UserID(c.userID)
	user, err := s.userRepository.GetById(ctx, userID)

	if err != nil {
		return errors.NewUserNotFoundError()
	}

	con, hmac, err := s.confirmationService.Create(ctx, userID)

	if err != nil {
		return errors.NewCannotCreateConfirmationError(err)
	}

	err = s.mailService.SendConfirmationMail(con.ID, user.Email, c.senderEmail, user.FullName(), *hmac)

	if err != nil {
		return err
	}

	return nil
}
