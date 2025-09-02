package confirm_user

import (
	"app/config/api"
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/application/services"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/port/errors"
	"app/internal/core/port/idp"
	"context"
)

type Handler struct {
	userRepository repositories.UserRepository
	idp            idp.IdentityProvider
	mailService    *services.MailService
	errors         errors.ErrorFactory
	senderEmail    string
}

func (h *Handler) Handle(ctx context.Context, c Command) error {
	usr, err := h.userRepository.GetById(ctx, user.ID(c.userID))

	if err != nil {
		return h.errors.New(errors.ErrValidation, "User not found", err)
	}

	if usr.IdPUserId == nil {
		return h.errors.New(errors.ErrValidation, "User is not linked to IdP", err)
	}

	err = h.idp.ConfirmUser(ctx, *usr.IdPUserId)

	if err != nil {
		return h.errors.New(errors.ErrValidation, "Error confirming user at IdP", err)
	}

	usr.Confirm()

	if err := h.userRepository.Update(ctx, usr); err != nil {
		return h.errors.New(errors.ErrDB, "Error updating user", err)
	}

	if err := h.mailService.SendUserConfirmedMail(usr.Email, user.Email(h.senderEmail), usr.FullName()); err != nil {
		return h.errors.New(errors.ErrMail, "Error sending user confirmed mail", err)
	}

	return nil
}

func NewHandler(
	userRepository repositories.UserRepository,
	idp idp.IdentityProvider,
	mailService *services.MailService,
	errors errors.ErrorFactory,
	// TODO dependency creep - this shouldn't be injected here! Use a port!
	cfg api.Config,
) *Handler {
	return &Handler{
		userRepository: userRepository,
		idp:            idp,
		mailService:    mailService,
		errors:         errors,
		senderEmail:    cfg.Mail.DefaultSender,
	}
}
