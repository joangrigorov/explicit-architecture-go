package presentation

import (
	"app/internal/presentation/web"
	home "app/internal/presentation/web/pages/home/controllers"
	id "app/internal/presentation/web/pages/identity/controllers"
	svc "app/internal/presentation/web/services"

	"go.uber.org/fx"
)

var pages = fx.Module("pages", fx.Provide(
	home.NewHome,
	id.NewSignUp,
	id.NewPasswordSetup,
))

var framework = fx.Module("framework", fx.Provide(
	web.NewTemplate,
), fx.Invoke(
	web.RegisterRoutes,
))

var services = fx.Module("services", fx.Provide(
	svc.NewActivityPlannerClient,
	svc.NewIdentityService,
))

var Web = fx.Module("presentation/web",
	framework,
	pages,
	services,
)
