package fx

import (
	"app/internal/presentation/api"
	"app/internal/presentation/api/component/activity/v1/controllers/activities"
	"app/internal/presentation/api/component/user/v1/controllers"

	"go.uber.org/fx"
)

var Presentation = fx.Module("presentation",
	fx.Module("api", fx.Provide(
		activities.NewController,
		controllers.NewRegistrationController,
	), fx.Invoke(
		api.RegisterRoutes,
	)),
)
