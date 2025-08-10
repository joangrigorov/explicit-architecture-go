package bootstrap

import (
	"app/config"
	"app/internal/infrastructure/framework/http"
	"app/internal/infrastructure/framework/validation"
	"app/internal/infrastructure/persistence/ent"
	"app/internal/infrastructure/persistence/ent/activity"
	"app/internal/infrastructure/persistence/ent/attendance"
	"app/internal/presentation/web/core/component/activity/v1/controllers/activities"

	"go.uber.org/fx"
)

var providers = fx.Options(
	fx.Provide(
		// configuration providers
		config.NewConfig,

		// database (ent) related providers
		ent.NewDB,
		activity.NewClient,
		attendance.NewClient,

		// framework providers
		http.NewGinEngine,
		http.NewRouter,
		validation.NewValidatorValidate,
		validation.NewTranslator,

		// repository adapter providers
		activity.NewActivityRepository,
		attendance.NewAttendanceRepository,

		// web controller providers
		activities.NewController,
	),
)
