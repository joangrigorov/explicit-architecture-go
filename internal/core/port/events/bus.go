package events

import "app/internal/core/shared_kernel/events"

type EventBus interface {
	Publish(events.Event) error
}
