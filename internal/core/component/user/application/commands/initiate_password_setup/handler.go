package initiate_password_setup

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/application/services"
	domain "app/internal/core/component/user/domain/user"
	"app/internal/core/port/errors"
	"context"
)

type Handler struct {
	userRepository      repositories.UserRepository
	confirmationService *services.ConfirmationService
	mailService         *services.MailService
	errors              errors.ErrorFactory
}

func NewHandler(
	userRepository repositories.UserRepository,
	confirmationService *services.ConfirmationService,
	mailService *services.MailService,
	errors errors.ErrorFactory,
) *Handler {
	return &Handler{
		userRepository:      userRepository,
		confirmationService: confirmationService,
		mailService:         mailService,
		errors:              errors,
	}
}

func (s *Handler) Handle(ctx context.Context, c Command) error {
	userID := domain.ID(c.userID)
	user, err := s.userRepository.GetById(ctx, userID)

	if err != nil {
		return s.errors.New(errors.ErrValidation, "User does not exist", err)
	}

	cfr, hmac, err := s.confirmationService.Create(ctx, userID)

	if err != nil {
		return s.errors.New(errors.ErrDB, "Creating confirmation failed", err)
	}

	err = s.mailService.SendPasswordSetupMail(cfr.ID, user.Email, domain.Email(c.senderEmail), user.FullName(), *hmac)

	if err != nil {
		return s.errors.New(errors.ErrExternal, "Sending email failed", err)
	}

	return nil
}
