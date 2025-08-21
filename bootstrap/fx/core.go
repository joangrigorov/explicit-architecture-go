package fx

import (
	"app/internal/core/component/user/application/queries"

	"go.uber.org/fx"
)

var Core = fx.Module("core",
	fx.Module("components",
		fx.Module("user", fx.Provide(
			queries.NewFindUserByIDHandler,
		)),
	),
)
