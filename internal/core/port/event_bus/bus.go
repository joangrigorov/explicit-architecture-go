package event_bus

import (
	"app/internal/core/shared_kernel/events"
	"context"
)

type EventBus interface {
	Publish(context.Context, events.Event) error
}
