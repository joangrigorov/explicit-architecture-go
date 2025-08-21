package bootstrap

import (
	d "app/bootstrap/fx"
	p "app/bootstrap/fx/presentation"

	"go.uber.org/fx"
)

func APIApp() *fx.App {
	return fx.New(
		d.Config,
		d.Subscribers,
		d.Infrastructure,
		d.Core,
		d.Migrations,
		p.Api,
		RunServer,
	)
}
