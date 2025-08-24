package commands

import (
	"app/internal/core/component/user/application/errors"
	"app/internal/core/component/user/application/repositories"
	eventBus "app/internal/core/port/event_bus"
	"app/internal/core/port/idp"
	"app/internal/core/shared_kernel/domain"
	"app/internal/core/shared_kernel/events"
	"context"
	"encoding/json"
)

type CreateIdPUserCommand struct {
	userID   domain.UserID
	username string
	email    string
	password string
}

func (c CreateIdPUserCommand) LogBody() ([]byte, error) {
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
	eventBus       eventBus.EventBus
}

func NewCreateIdPUserHandler(
	userRepository repositories.UserRepository,
	idp idp.IdentityProvider,
	eventBus eventBus.EventBus,
) *CreateIdPUserHandler {
	return &CreateIdPUserHandler{
		userRepository: userRepository,
		idp:            idp,
		eventBus:       eventBus,
	}
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

	err = h.userRepository.Update(ctx, user)

	if err != nil {
		return err
	}

	return h.eventBus.Publish(ctx, events.NewIdPUserCreated(c.userID, *idpUserID))
}
