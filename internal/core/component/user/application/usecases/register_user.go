package usecases

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain"
	"app/internal/core/port/event_bus"
	"app/internal/core/port/uuid"
	"app/internal/core/shared_kernel/events"
	"context"
)

type RegisterUser struct {
	userRepository repositories.UserRepository
	uuidGenerator  uuid.Generator
	eventBus       event_bus.EventBus
}

func NewRegisterUser(
	ur repositories.UserRepository,
	uuidGenerator uuid.Generator,
	eventBus event_bus.EventBus,
) *RegisterUser {
	return &RegisterUser{
		userRepository: ur,
		uuidGenerator:  uuidGenerator,
		eventBus:       eventBus,
	}
}

func (r *RegisterUser) Execute(
	ctx context.Context,
	username string,
	password string,
	email string,
	firstName string,
	lastName string,
) error {
	id := events.UserId(r.uuidGenerator.Generate())
	user := domain.NewUser(id, username, email, firstName, lastName, &domain.Member{})

	err := r.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	return r.eventBus.Publish(&events.UserCreated{
		UserId:   id,
		Email:    email,
		Password: password,
	})
}
