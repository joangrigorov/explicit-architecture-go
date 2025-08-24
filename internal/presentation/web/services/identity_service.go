package services

import (
	"app/config/web"
	"fmt"
	"net/url"
)

type IdentityService struct {
	keycloakCfg web.Keycloak
}

func (s IdentityService) SignInURL(state string, nonce string) string {
	authURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/auth", s.keycloakCfg.Host, s.keycloakCfg.Realm)

	params := url.Values{}
	params.Add("client_id", s.keycloakCfg.ClientId)
	params.Add("redirect_uri", s.keycloakCfg.RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", "openid profile email")
	params.Add("state", state)
	params.Add("nonce", nonce)

	return authURL + "?" + params.Encode()
}

func NewIdentityService(cfg *web.Config) *IdentityService {
	return &IdentityService{
		keycloakCfg: cfg.Keycloak,
	}
}
