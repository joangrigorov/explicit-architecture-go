package events

import (
	userSubscribers "app/internal/core/component/user/application/subscribers"
	"app/internal/core/component/user/domain/user"
	eventBus "app/internal/infrastructure/framework/events"
)

func WireSubscribers(
	eventBus *eventBus.SimpleEventBus,

	sendPsswdSetupMailSub *userSubscribers.SendSetPasswordMailSubscriber,
	confirmUserSub *userSubscribers.ConfirmUserSubscriber,
) {
	eventBus.Subscribe(sendPsswdSetupMailSub, user.CreatedEvent{})
	eventBus.Subscribe(confirmUserSub, user.IdPUserLinkedEvent{})
}
