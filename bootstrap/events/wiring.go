package events

import (
	"app/bootstrap/events/subscribers"
	"app/internal/core/shared_kernel/events"
	eventBus "app/internal/infrastructure/event_bus"
)

func WireSubscribers(
	eventBus *eventBus.SimpleEventBus,
	kcUserSub *subscribers.CreateKeycloakUserSubscriber,
) {
	eventBus.Subscribe(kcUserSub, events.UserCreated{})
}
