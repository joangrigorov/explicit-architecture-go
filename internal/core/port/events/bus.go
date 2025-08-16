package events

import events2 "app/internal/core/shared_kernel/events"

type EventBus interface {
	Publish(events2.Event)
}
