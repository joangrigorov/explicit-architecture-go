package events

import (
	"app/internal/core/shared_kernel/events"
	"app/internal/infrastructure/component/user/subscribers"
	eventBus "app/internal/infrastructure/framework/event_bus"
)

func WireSubscribers(
	eventBus *eventBus.SimpleEventBus,
	kcUserSub *subscribers.CreateKeycloakUserSubscriber,
) {
	eventBus.Subscribe(kcUserSub, events.UserCreated{})
}
