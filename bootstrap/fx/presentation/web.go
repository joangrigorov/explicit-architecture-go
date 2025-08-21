package presentation

import (
	"app/internal/presentation/web"
	"app/internal/presentation/web/pages/home/controllers"

	"go.uber.org/fx"
)

var Web = fx.Module("presentation/web", fx.Provide(
	controllers.NewHome,
), fx.Invoke(
	web.RegisterRoutes,
))
