package event_bus

import (
	eventBus "app/internal/core/shared_kernel/events"
	"context"
)

type Subscriber interface {
	Dispatch(context.Context, eventBus.Event) error
}
