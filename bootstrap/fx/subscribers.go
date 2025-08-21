package fx

import (
	"app/internal/infrastructure/component/user/subscribers"

	"go.uber.org/fx"
)

var Subscribers = fx.Module("subscribers", fx.Provide(
	subscribers.NewCreateKeycloakUserSubscriber,
))
