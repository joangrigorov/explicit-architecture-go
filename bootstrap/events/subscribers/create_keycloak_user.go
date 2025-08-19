package subscribers

import (
	. "app/internal/core/component/user/application/commands"
	"app/internal/core/port/cqrs"
	"app/internal/core/port/logging"
	"app/internal/core/shared_kernel/events"
	"context"
	"errors"
	"fmt"
	"reflect"

	"go.opentelemetry.io/otel/trace"
)

type CreateKeycloakUserSubscriber struct {
	commandBus cqrs.CommandBus
	logger     logging.Logger
	tracer     trace.Tracer
}

func NewCreateKeycloakUserSubscriber(
	commandBus cqrs.CommandBus,
	logger logging.Logger,
	tracer trace.Tracer,
) *CreateKeycloakUserSubscriber {
	return &CreateKeycloakUserSubscriber{commandBus: commandBus, logger: logger, tracer: tracer}
}

func (c *CreateKeycloakUserSubscriber) Dispatch(ctx context.Context, e events.Event) error {
	ctx, span := c.tracer.Start(ctx, fmt.Sprintf("Event %T handled by %T", e, c))
	defer span.End()
	c.logger.Debug(fmt.Sprintf("Entered the %T subscriber", c))
	defer c.logger.Debug(fmt.Sprintf("Exit the %T subscriber", c))

	event, ok := e.(events.UserCreated)
	if !ok {
		return errors.New(fmt.Sprintf("%T subscriber cannot subscribe to %s", c, reflect.TypeOf(e).Name()))
	}

	return c.commandBus.Dispatch(context.WithoutCancel(ctx), NewCreateIdPUserCommand(
		event.UserId(),
		event.Username(),
		event.Email(),
		event.Password(),
	))
}
