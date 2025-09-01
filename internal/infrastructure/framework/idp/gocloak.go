package idp

import (
	"app/config/api"

	"github.com/Nerzal/gocloak/v13"
)

func NewGoCloakClient(cfg api.Config) *gocloak.GoCloak {
	return gocloak.NewClient(cfg.Keycloak.Url)
}
