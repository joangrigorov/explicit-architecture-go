package commands

import (
	"app/internal/core/component/user/application/errors"
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/port/idp"
	"app/internal/core/shared_kernel/domain"
	"context"
	"encoding/json"
)

type CreateIdPUserCommand struct {
	userID   domain.UserID
	username string
	email    string
	password string
}

func (c CreateIdPUserCommand) Serialize() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userId":   c.userID,
		"username": c.username,
		"email":    c.email,
	})
}

func NewCreateIdPUserCommand(
	userID domain.UserID,
	username string,
	email string,
	password string,
) CreateIdPUserCommand {
	return CreateIdPUserCommand{
		userID:   userID,
		username: username,
		email:    email,
		password: password,
	}
}

type CreateIdPUserHandler struct {
	userRepository repositories.UserRepository
	idp            idp.IdentityProvider
}

func NewCreateIdPUserHandler(
	userRepository repositories.UserRepository,
	idp idp.IdentityProvider,
) *CreateIdPUserHandler {
	return &CreateIdPUserHandler{userRepository: userRepository, idp: idp}
}

func (h *CreateIdPUserHandler) Handle(ctx context.Context, c CreateIdPUserCommand) error {
	user, err := h.userRepository.GetById(ctx, c.userID)

	if err != nil {
		return errors.NewUserNotFoundError()
	}

	idpUserID, err := h.idp.CreateUser(ctx, c.userID, c.username, c.email, c.password)

	if err != nil {
		return errors.NewCannotCreateIdPUserError(err)
	}

	user.IdPUserId = idpUserID

	return h.userRepository.Update(ctx, user)
}
