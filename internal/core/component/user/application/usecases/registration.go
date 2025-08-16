package usecases

import (
	. "app/internal/core/component/user/application/repositories"
	. "app/internal/core/component/user/domain"
	. "app/internal/core/port/events"
	"app/internal/core/port/uuid"
	. "app/internal/core/shared_kernel/domain"
	. "app/internal/core/shared_kernel/events"
	"context"
)

type Registration struct {
	userRepository UserRepository
	uuidGenerator  uuid.Generator
	eventBus       EventBus
}

func NewRegistration(
	ur UserRepository,
	uuidGenerator uuid.Generator,
	eventBus EventBus,
) *Registration {
	return &Registration{
		userRepository: ur,
		uuidGenerator:  uuidGenerator,
		eventBus:       eventBus,
	}
}

func (r *Registration) Execute(
	ctx context.Context,
	username string,
	password string,
	email string,
	firstName string,
	lastName string,
) error {
	id := UserID(r.uuidGenerator.Generate())
	user := NewUser(id, username, email, firstName, lastName, &Member{})

	err := r.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	r.eventBus.Publish(NewUserCreated(id, email, password))

	return nil
}
