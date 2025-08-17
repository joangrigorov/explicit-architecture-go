package commands

import (
	"app/internal/core/component/user/application/repositories"
	. "app/internal/core/component/user/domain"
	eventBus "app/internal/core/port/events"
	"app/internal/core/port/uuid"
	. "app/internal/core/shared_kernel/domain"
	"app/internal/core/shared_kernel/events"
	"context"
)

type RegisterUserCommand struct {
	username  string
	password  string
	email     string
	firstName string
	lastName  string
}

func NewRegisterUserCommand(
	username string,
	password string,
	email string,
	firstName string,
	lastName string,
) RegisterUserCommand {
	return RegisterUserCommand{
		username:  username,
		password:  password,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
	}
}

type RegisterUserCommandHandler struct {
	userRepository repositories.UserRepository
	eventBus       eventBus.EventBus
	uuidGenerator  uuid.Generator
}

func NewRegisterUserCommandHandler(
	userRepository repositories.UserRepository,
	eventBus eventBus.EventBus,
	uuidGenerator uuid.Generator,
) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{userRepository: userRepository, eventBus: eventBus, uuidGenerator: uuidGenerator}
}

func (h *RegisterUserCommandHandler) Handle(ctx context.Context, c RegisterUserCommand) error {
	id := UserID(h.uuidGenerator.Generate())
	user := NewUser(id, c.username, c.email, c.firstName, c.lastName, &Member{})

	if err := h.userRepository.Create(ctx, user); err != nil {
		return err
	}

	return h.eventBus.Publish(events.NewUserCreated(id, c.username, c.email, c.password))
}
