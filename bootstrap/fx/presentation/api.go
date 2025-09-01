package presentation

import (
	"app/internal/presentation/api"
	"app/internal/presentation/api/component/activity/v1/controllers/activities"
	"app/internal/presentation/api/component/user/v1/controllers"
	"app/internal/presentation/api/shared/errors"

	"go.uber.org/fx"
)

var Api = fx.Module("presentation/api", fx.Provide(
	activities.NewController,
	controllers.NewRegistration,
	controllers.NewVerification,
	controllers.NewMe,
	errors.NewHandler,
), fx.Invoke(
	api.RegisterRoutes,
))
