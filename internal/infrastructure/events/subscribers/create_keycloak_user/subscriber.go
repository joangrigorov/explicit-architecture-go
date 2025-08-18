package create_keycloak_user

import (
	. "app/internal/core/component/user/application/commands"
	"app/internal/core/port/commands"
	"app/internal/core/port/logging"
	"app/internal/core/shared_kernel/events"
	eventBus "app/internal/infrastructure/events"
	"context"
	"errors"
	"fmt"
	"reflect"

	"go.opentelemetry.io/otel/trace"
)

type Subscriber struct {
	commandBus commands.CommandBus
	logger     logging.Logger
	tracer     trace.Tracer
}

func (c *Subscriber) Dispatch(ctx context.Context, e events.Event) error {
	ctx, span := c.tracer.Start(ctx, fmt.Sprintf("Event %T handled by %T", e, c))
	defer span.End()
	c.logger.Debug(fmt.Sprintf("Entered the %T subscriber", c))
	defer c.logger.Debug(fmt.Sprintf("Exit the %T subscriber", c))

	event, ok := e.(events.UserCreated)
	if !ok {
		return errors.New(fmt.Sprintf("create_keycloak_user subscriber cannot subscribe to %s", reflect.TypeOf(e).Name()))
	}

	return c.commandBus.Dispatch(context.WithoutCancel(ctx), NewCreateIdPUserCommand(
		event.UserId(),
		event.Username(),
		event.Email(),
		event.Password(),
	))
}

func Register(
	eventBus *eventBus.SimpleEventBus,
	commandBus commands.CommandBus,
	logger logging.Logger,
	tracer trace.Tracer,
) {
	subscriber := &Subscriber{commandBus: commandBus, logger: logger, tracer: tracer}
	eventBus.Subscribe(subscriber, events.UserCreated{})
}
