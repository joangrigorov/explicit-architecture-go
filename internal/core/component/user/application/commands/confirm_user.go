package commands

import (
	"app/internal/core/component/user/application/errors"
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/port/idp"
	"app/internal/core/shared_kernel/domain"
	"context"
	"encoding/json"
)

type ConfirmUserCommand struct {
	userID string
}

func (c ConfirmUserCommand) Serialize() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userID": c.userID,
	})
}

func NewConfirmUserCommand(userID string) ConfirmUserCommand {
	return ConfirmUserCommand{userID: userID}
}

type ConfirmUserCommandHandler struct {
	userRepository repositories.UserRepository
	idp            idp.IdentityProvider
}

func NewConfirmUserCommandHandler(
	userRepository repositories.UserRepository,
	idp idp.IdentityProvider,
) *ConfirmUserCommandHandler {
	return &ConfirmUserCommandHandler{userRepository: userRepository, idp: idp}
}

func (h *ConfirmUserCommandHandler) Handle(ctx context.Context, c ConfirmUserCommand) error {
	user, err := h.userRepository.GetById(ctx, domain.UserID(c.userID))

	if err != nil {
		return errors.NewUserNotFoundError()
	}

	if user.IdPUserId == nil {
		return errors.NewIdPUserNotConnectedError()
	}

	err = h.idp.ConfirmUser(ctx, *user.IdPUserId)

	if err != nil {
		return errors.NewIdPRequestError(err)
	}

	user.Confirm()

	return h.userRepository.Update(ctx, user)
}
