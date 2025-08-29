package subscribers

import (
	"app/internal/core/component/user/application/commands/confirm_user"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/port/cqrs"
	"app/internal/core/port/errors"
	"app/internal/core/port/events"
	"context"
)

type ConfirmUserSubscriber struct {
	commandBus cqrs.CommandBus
	errors     errors.ErrorFactory
}

func NewConfirmUserSubscriber(commandBus cqrs.CommandBus, errors errors.ErrorFactory) *ConfirmUserSubscriber {
	return &ConfirmUserSubscriber{commandBus: commandBus, errors: errors}
}

func (s *ConfirmUserSubscriber) Dispatch(ctx context.Context, event events.Event) error {
	e, ok := event.(user.IdPUserLinkedEvent)
	if !ok {
		return s.errors.New(errors.ErrTypeMismatch, "Subscriber listening to wrong event type", nil)
	}

	command := confirm_user.NewCommand(e.UserID().String())
	if err := s.commandBus.Dispatch(ctx, command); err != nil {
		return s.errors.New(errors.ErrCommandHandlingError, "Error handling command", err)
	}

	return nil
}
