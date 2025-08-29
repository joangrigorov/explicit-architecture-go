package bootstrap

import (
	d "app/bootstrap/fx"
	p "app/bootstrap/fx/presentation"
	"app/config/api"

	"go.uber.org/fx"
)

func APIApp() *fx.App {
	return fx.New(
		fx.Module("config", fx.Provide(api.NewConfig)),
		d.Infrastructure,
		d.Core,
		d.Migrations,
		p.Api,
		RunServer,
	)
}
