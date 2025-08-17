package create_keycloak_user

import (
	"app/internal/core/component/user/application/usecases"
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
	createIdPUser *usecases.CreateIdPUser
	logger        logging.Logger
	tracer        trace.Tracer
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

	return c.createIdPUser.Execute(
		context.Background(), // we switch the context intentionally so that existing transactions are not impacted
		event.UserId(),
		event.Username(),
		event.Email(),
		event.Password(),
	)
}

func Register(
	bus *eventBus.SimpleEventBus,
	createIdPUser *usecases.CreateIdPUser,
	logger logging.Logger,
	tracer trace.Tracer,
) {
	subscriber := &Subscriber{createIdPUser: createIdPUser, logger: logger, tracer: tracer}
	bus.Subscribe(subscriber, events.UserCreated{})
}
