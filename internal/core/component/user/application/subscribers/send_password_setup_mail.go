package subscribers

import (
	"app/config/api"
	"app/internal/core/component/user/application/commands/initiate_password_setup"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/port/cqrs"
	"app/internal/core/port/errors"
	"app/internal/core/port/events"
	"context"
)

type SendSetPasswordMailSubscriber struct {
	commandBus  cqrs.CommandBus
	senderEmail string
	errors      errors.ErrorFactory
}

func (s *SendSetPasswordMailSubscriber) Dispatch(ctx context.Context, event events.Event) error {
	e, ok := event.(user.CreatedEvent)
	if !ok {
		return s.errors.New(errors.ErrTypeMismatch, "Subscriber listening to wrong event type", nil)
	}

	command := initiate_password_setup.NewCommand(e.UserID().String(), s.senderEmail)
	if err := s.commandBus.Dispatch(ctx, command); err != nil {
		return s.errors.New(errors.ErrCommandHandling, "Error handling command", err)
	}

	return nil
}

func NewSendSetPasswordMailSubscriber(
	commandBus cqrs.CommandBus,
	errors errors.ErrorFactory,
	// TODO dependency creep - this shouldn't be injected here! Use a port!
	cfg api.Config,
) *SendSetPasswordMailSubscriber {
	return &SendSetPasswordMailSubscriber{
		commandBus:  commandBus,
		senderEmail: cfg.Mail.DefaultSender,
		errors:      errors,
	}
}
