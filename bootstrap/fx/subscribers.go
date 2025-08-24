package fx

import (
	userSubs "app/internal/infrastructure/component/user/subscribers"

	"go.uber.org/fx"
)

var Subscribers = fx.Module("subscribers", fx.Provide(
	userSubs.NewCreateKeycloakUserSubscriber,
	userSubs.NewSendConfirmationMailSubscriber,
))
