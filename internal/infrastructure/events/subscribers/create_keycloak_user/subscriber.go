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
)

type Subscriber struct {
	createIdPUser *usecases.CreateIdPUser
	logger        logging.Logger
}

func (c *Subscriber) Dispatch(e events.Event) error {
	c.logger.Debug("Entered the create_keycloak_user subscriber")
	defer c.logger.Debug("Exit the create_keycloak_user subscriber")
	
	event, ok := e.(events.UserCreated)
	if !ok {
		return errors.New(fmt.Sprintf("create_keycloak_user subscriber cannot subscribe to %s", reflect.TypeOf(e).Name()))
	}

	return c.createIdPUser.Execute(
		context.Background(),
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
) {
	subscriber := &Subscriber{createIdPUser: createIdPUser, logger: logger}
	bus.Subscribe(subscriber, events.UserCreated{})
}
