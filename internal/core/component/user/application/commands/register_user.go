package commands

import (
	"app/internal/core/component/user/application/repositories"
	. "app/internal/core/component/user/domain"
	eventBus "app/internal/core/port/events"
	. "app/internal/core/shared_kernel/domain"
	"app/internal/core/shared_kernel/events"
	"context"
	"encoding/json"
)

type RegisterUserCommand struct {
	userID    string
	username  string
	password  string
	email     string
	firstName string
	lastName  string
}

func (r RegisterUserCommand) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userID":    r.userID,
		"username":  r.username,
		"email":     r.email,
		"firstName": r.firstName,
		"lastName":  r.lastName,
	})
}

func NewRegisterUserCommand(
	userID string,
	username string,
	password string,
	email string,
	firstName string,
	lastName string,
) RegisterUserCommand {
	return RegisterUserCommand{
		userID:    userID,
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
}

func NewRegisterUserCommandHandler(
	userRepository repositories.UserRepository,
	eventBus eventBus.EventBus,
) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{userRepository: userRepository, eventBus: eventBus}
}

func (h *RegisterUserCommandHandler) Handle(ctx context.Context, c RegisterUserCommand) error {
	id := UserID(c.userID)
	user := NewUser(id, c.username, c.email, c.firstName, c.lastName, &Member{})

	if err := h.userRepository.Create(ctx, user); err != nil {
		return err
	}

	return h.eventBus.Publish(ctx, events.NewUserCreated(id, c.username, c.email, c.password))
}
