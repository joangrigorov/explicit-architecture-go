package identity

import (
	"app/config/web"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type IdPService struct {
	keycloakCfg web.Keycloak
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

func (s *IdPService) SignInURL(state string, nonce string) string {
	authURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/auth", s.keycloakCfg.UIHost, s.keycloakCfg.Realm)

	params := url.Values{}
	params.Add("client_id", s.keycloakCfg.ClientId)
	params.Add("redirect_uri", s.keycloakCfg.RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", "openid profile email")
	params.Add("state", state)
	params.Add("nonce", nonce)

	return authURL + "?" + params.Encode()
}

func (s *IdPService) ExchangeCodeForToken(code string) (*TokenResponse, error) {
	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token",
		s.keycloakCfg.APIHost, s.keycloakCfg.Realm)

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", s.keycloakCfg.ClientId)
	if s.keycloakCfg.ClientSecret != "" {
		data.Set("client_secret", s.keycloakCfg.ClientSecret)
	}
	data.Set("redirect_uri", s.keycloakCfg.RedirectURI)

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token request failed: %s", resp.Status)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func (s *IdPService) RefreshToken(refreshToken string) (*TokenResponse, error) {
	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token",
		s.keycloakCfg.APIHost, s.keycloakCfg.Realm)

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", s.keycloakCfg.ClientId)
	if s.keycloakCfg.ClientSecret != "" {
		data.Set("client_secret", s.keycloakCfg.ClientSecret)
	}

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("refresh token request failed: %s", resp.Status)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func (s *IdPService) Logout(refreshToken string) error {
	logoutURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/logout",
		s.keycloakCfg.APIHost, s.keycloakCfg.Realm)

	data := url.Values{}
	data.Set("client_id", s.keycloakCfg.ClientId)
	if s.keycloakCfg.ClientSecret != "" {
		data.Set("client_secret", s.keycloakCfg.ClientSecret)
	}
	log.Println(refreshToken)
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", logoutURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyStr := string(bodyBytes)

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("logout request failed: %s; %s", resp.Status, bodyStr)
	}

	return nil
}

func NewIdentityService(cfg web.Config) *IdPService {
	return &IdPService{
		keycloakCfg: cfg.Keycloak,
	}
}
