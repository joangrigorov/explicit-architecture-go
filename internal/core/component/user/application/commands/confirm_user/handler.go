package confirm_user

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/port/errors"
	"app/internal/core/port/idp"
	"context"
)

type Handler struct {
	userRepository repositories.UserRepository
	idp            idp.IdentityProvider
	errors         errors.ErrorFactory
}

func NewHandler(
	userRepository repositories.UserRepository,
	idp idp.IdentityProvider,
	errors errors.ErrorFactory,
) *Handler {
	return &Handler{userRepository: userRepository, idp: idp, errors: errors}
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

	err = h.userRepository.Update(ctx, usr)

	if err != nil {
		return h.errors.New(errors.ErrDB, "Error updating user", err)
	}

	return nil
}
