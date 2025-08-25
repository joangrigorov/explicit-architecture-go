package commands

import (
	"app/internal/core/component/user/application/repositories"
	. "app/internal/core/component/user/domain"
	"app/internal/core/port/errors"
	eventBus "app/internal/core/port/event_bus"
	. "app/internal/core/shared_kernel/domain"
	"app/internal/core/shared_kernel/events"
	"context"
)

type RegisterUserCommandHandler struct {
	userRepository repositories.UserRepository
	eventBus       eventBus.EventBus
	errors         errors.ErrorFactory
}

func NewRegisterUserCommandHandler(
	userRepository repositories.UserRepository,
	eventBus eventBus.EventBus,
	errors errors.ErrorFactory,
) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{
		userRepository: userRepository,
		eventBus:       eventBus,
		errors:         errors,
	}
}

func (h *RegisterUserCommandHandler) Handle(ctx context.Context, c RegisterUserCommand) error {
	id := UserID(c.userID)

	if user, err := h.userRepository.GetById(ctx, id); err == nil && user != nil {
		return h.errors.New(errors.ErrConflict, "User already exists", nil)
	}

	user := NewUser(id, c.username, c.email, c.firstName, c.lastName, &Member{})

	if err := h.userRepository.Create(ctx, user); err != nil {
		return h.errors.New(errors.ErrDB, "Error creating user", err)
	}

	return h.eventBus.Publish(ctx, events.NewUserCreated(id, c.username, c.email, c.password))
}
