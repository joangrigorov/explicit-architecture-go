package complete_password_setup

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/port/errors"
	eventBus "app/internal/core/port/events"
	"app/internal/core/port/idp"
	"context"
)

type Handler struct {
	userRepository repositories.UserRepository
	idp            idp.IdentityProvider
	eventBus       eventBus.EventBus
	errors         errors.ErrorFactory
}

func NewHandler(
	userRepository repositories.UserRepository,
	idp idp.IdentityProvider,
	eventBus eventBus.EventBus,
	errors errors.ErrorFactory,
) *Handler {
	return &Handler{
		userRepository: userRepository,
		idp:            idp,
		eventBus:       eventBus,
		errors:         errors,
	}
}

func (h *Handler) Handle(ctx context.Context, c Command) error {
	userID := user.ID(c.userID)
	usr, err := h.userRepository.GetById(ctx, userID)

	if err != nil {
		return h.errors.New(errors.ErrValidation, "User not found", err)
	}

	if usr.IdPUserId != nil {
		return h.errors.New(errors.ErrValidation, "User is already linked to IdP", nil)
	}

	idpUserID, err := h.idp.CreateUser(ctx, userID, usr.Username.String(), usr.Email.String(), c.password)

	if err != nil {
		return h.errors.New(errors.ErrDB, "Cannot create user in IdP", err)
	}

	usr.LinkToIdP(*idpUserID)

	err = h.userRepository.Update(ctx, usr)

	if err != nil {
		return h.errors.New(errors.ErrDB, "Error updating user", err)
	}

	return nil
}
