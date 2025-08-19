package fx

import (
	"app/bootstrap/events/subscribers"

	"go.uber.org/fx"
)

var Subscribers = fx.Module("subscribers", fx.Provide(
	subscribers.NewCreateKeycloakUserSubscriber,
))
