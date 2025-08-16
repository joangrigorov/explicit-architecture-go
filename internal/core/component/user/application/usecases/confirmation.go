package usecases

import (
	"app/internal/core/component/user/application/errors"
	"app/internal/core/component/user/application/repositories"
	eventBus "app/internal/core/port/events"
	"app/internal/core/shared_kernel/domain"
	"app/internal/core/shared_kernel/events"
	"context"
)

type Confirmation struct {
	userRepository repositories.UserRepository
	eventBus       eventBus.EventBus
}

func NewConfirmation(userRepository repositories.UserRepository, eventBus eventBus.EventBus) *Confirmation {
	return &Confirmation{
		userRepository: userRepository,
		eventBus:       eventBus,
	}
}

func (uc *Confirmation) Execute(ctx context.Context, userID string) error {
	user, err := uc.userRepository.GetById(ctx, domain.UserID(userID))

	if err != nil {
		return errors.NewUserNotFoundError()
	}

	user.Confirm()

	err = uc.userRepository.Update(ctx, user)

	if err != nil {
		return err
	}

	uc.eventBus.Publish(events.NewUserConfirmed(user.ID))

	return err
}
