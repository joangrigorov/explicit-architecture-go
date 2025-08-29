package fx

import (
	"app/bootstrap/fx/core/component"

	"go.uber.org/fx"
)

var Core = fx.Module("core",
	fx.Module("component",
		component.User,
	),
)
