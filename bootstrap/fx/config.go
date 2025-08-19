package fx

import (
	"app/config"

	"go.uber.org/fx"
)

var Config = fx.Module("config", fx.Provide(config.NewConfig))
