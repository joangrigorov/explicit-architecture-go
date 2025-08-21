package idp

import (
	"app/config"

	"github.com/Nerzal/gocloak/v13"
)

func NewGoCloakClient(cfg *config.Config) *gocloak.GoCloak {
	return gocloak.NewClient(cfg.Keycloak.Url)
}
