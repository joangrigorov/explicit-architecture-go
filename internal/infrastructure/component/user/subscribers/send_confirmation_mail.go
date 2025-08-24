package subscribers

import (
	"app/config/api"
	"app/internal/core/component/user/application/commands"
	"app/internal/core/port/cqrs"
	"app/internal/core/port/logging"
	"app/internal/core/shared_kernel/events"
	"context"
	"errors"
	"fmt"
	"reflect"

	"go.opentelemetry.io/otel/trace"
)

type SendConfirmationMailSubscriber struct {
	commandBus  cqrs.CommandBus
	logger      logging.Logger
	tracer      trace.Tracer
	senderEmail string
}

func NewSendConfirmationMailSubscriber(
	commandBus cqrs.CommandBus,
	cfg *api.Config,
	logger logging.Logger,
	tracer trace.Tracer,
) *SendConfirmationMailSubscriber {
	return &SendConfirmationMailSubscriber{
		commandBus:  commandBus,
		senderEmail: cfg.Mail.DefaultSender,
		logger:      logger,
		tracer:      tracer,
	}
}

func (s *SendConfirmationMailSubscriber) Dispatch(ctx context.Context, e events.Event) error {
	ctx, span := s.tracer.Start(ctx, fmt.Sprintf("Event %T handled by %T", e, s))
	defer span.End()
	s.logger.Debug(fmt.Sprintf("Entered the %T subscriber", s))
	defer s.logger.Debug(fmt.Sprintf("Exit the %T subscriber", s))

	event, ok := e.(events.IdPUserCreated)
	if !ok {
		return errors.New(fmt.Sprintf("%T subscriber cannot subscribe to %s", s, reflect.TypeOf(e).Name()))
	}

	return s.commandBus.Dispatch(ctx, commands.NewSendConfirmationMailCommand(event.UserID(), s.senderEmail))
}
