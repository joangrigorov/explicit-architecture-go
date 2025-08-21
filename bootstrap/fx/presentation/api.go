package presentation

import (
	"app/internal/presentation/api"

	"go.uber.org/fx"
)

var Api = fx.Module("presentation/api", fx.Provide(
// TODO register controllers
), fx.Invoke(
	api.RegisterRoutes,
))
