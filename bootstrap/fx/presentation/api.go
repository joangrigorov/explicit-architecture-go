package presentation

import (
	"app/internal/presentation/api"
	"app/internal/presentation/api/component/activity/v1/controllers/activities"
	"app/internal/presentation/api/component/user/v1/controllers"

	"go.uber.org/fx"
)

var Api = fx.Module("presentation/api", fx.Provide(
	activities.NewController,
	controllers.NewRegistrationController,
	controllers.NewVerification,
), fx.Invoke(
	api.RegisterRoutes,
))
