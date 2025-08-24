package fx

import (
	ent2 "app/internal/infrastructure/component/activity/persistence/ent"
	ent3 "app/internal/infrastructure/component/attendance/persistence/ent"
	"app/internal/infrastructure/component/user/persistence/ent"

	"go.uber.org/fx"
)

var Migrations = fx.Module("database/migrations",
	fx.Invoke(
		ent2.MigrateSchema,
		ent3.MigrateSchema,
		ent.MigrateSchema,
	),
)
