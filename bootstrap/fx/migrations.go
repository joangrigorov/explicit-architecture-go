package fx

import (
	"app/database/ent/schema/activity"
	"app/database/ent/schema/attendance"
	"app/database/ent/schema/user"

	"go.uber.org/fx"
)

var Migrations = fx.Module("database/migrations",
	fx.Invoke(
		activity.MigrateSchema,
		attendance.MigrateSchema,
		user.MigrateSchema,
	),
)
