package presentation

import (
	"app/internal/infrastructure/framework/validation"
	"app/internal/presentation/web"
	home "app/internal/presentation/web/pages/home/controllers"
	id "app/internal/presentation/web/pages/identity/controllers"
	"app/internal/presentation/web/services/activity_planner"
	"app/internal/presentation/web/services/identity"
	"app/internal/presentation/web/services/session"

	"go.uber.org/fx"
)

var pages = fx.Module("pages", fx.Provide(
	home.NewHome,
	id.NewSignUp,
	id.NewSignOut,
	id.NewPasswordSetup,
	id.NewOAuth2,
))

var framework = fx.Module("framework", fx.Provide(
	web.NewTemplate,
	validation.NewValidatorValidate,
	validation.NewTranslator,
), fx.Invoke(
	web.RegisterRoutes,
	validation.RegisterRules,
))

var services = fx.Module("services", fx.Provide(
	activity_planner.NewClient,
	identity.NewIdentityService,
	session.NewSessionStore,
	session.NewFlash,
	identity.NewAuthenticationService,
))

var Web = fx.Module("presentation/web",
	framework,
	pages,
	services,
)
