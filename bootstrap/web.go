package bootstrap

import (
	d "app/bootstrap/fx"
	p "app/bootstrap/fx/presentation"

	"go.uber.org/fx"
)

func WebApp() *fx.App {
	return fx.New(
		d.Config,
		d.Http,
		d.Logging,
		p.Web,
		RunServer,
	)
}
