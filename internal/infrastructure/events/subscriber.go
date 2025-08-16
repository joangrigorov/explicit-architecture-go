package events

import (
	eventBus "app/internal/core/shared_kernel/events"
)

type Subscriber interface {
	Dispatch(eventBus.Event) error
}
