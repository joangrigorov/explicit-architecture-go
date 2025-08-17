package usecases

import (
	"app/internal/core/component/user/application"
	. "app/internal/core/component/user/domain"
	"app/internal/core/port/uuid"
	. "app/internal/core/shared_kernel/domain"
	"app/internal/core/shared_kernel/events"
	"context"
)

type Registration struct {
	uow           application.UnitOfWork
	uuidGenerator uuid.Generator
}

func NewRegistration(
	uow application.UnitOfWork,
	uuidGenerator uuid.Generator,
) *Registration {
	return &Registration{
		uow:           uow,
		uuidGenerator: uuidGenerator,
	}
}

func (r *Registration) Execute(
	ctx context.Context,
	username string,
	password string,
	email string,
	firstName string,
	lastName string,
) (*User, error) {
	id := UserID(r.uuidGenerator.Generate())
	user := NewUser(id, username, email, firstName, lastName, &Member{})

	err := r.uow.Do(ctx, func(tx application.UnitOfWorkTx) error {
		if err := tx.UserRepository().Create(ctx, user); err != nil {
			return err
		}

		return tx.EventBus().Publish(events.NewUserCreated(id, username, email, password))
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
