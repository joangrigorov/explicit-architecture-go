package idp

import (
	"app/config/api"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/port/idp"
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/Nerzal/gocloak/v13"
)

type KeycloakIdentityProvider struct {
	client       *gocloak.GoCloak
	clientID     string
	clientSecret string
	scopes       string
	realm        string

	mu    sync.Mutex
	token *gocloak.JWT
}

func NewKeycloakIdentityProvider(client *gocloak.GoCloak, cfg api.Config) idp.IdentityProvider {
	return &KeycloakIdentityProvider{
		client:       client,
		clientID:     cfg.Keycloak.ClientId,
		clientSecret: cfg.Keycloak.ClientSecret,
		scopes:       cfg.Keycloak.Scopes,
		realm:        cfg.Keycloak.Realm,
	}
}

func (i *KeycloakIdentityProvider) getToken(ctx context.Context) (*gocloak.JWT, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.token != nil && i.token.ExpiresIn > 0 && time.Now().Before(time.Unix(int64(i.token.ExpiresIn), 0)) {
		return i.token, nil
	}

	token, err := i.client.LoginClient(ctx, i.clientID, i.clientSecret, i.realm)
	if err != nil {
		return nil, err
	}

	i.token = token

	return token, nil
}

func (i *KeycloakIdentityProvider) CreateUser(
	ctx context.Context,
	userID user.ID,
	username string,
	email string,
	password string,
) (*user.IdPUserID, error) {
	token, err := i.getToken(ctx)
	if err != nil {
		return nil, err
	}

	usr := gocloak.User{
		Email:    gocloak.StringP(email),
		Username: gocloak.StringP(username),
		Enabled:  gocloak.BoolP(false),
		Attributes: &map[string][]string{
			"app_user_id": {userID.String()},
		},
	}

	id, err := i.client.CreateUser(ctx, token.AccessToken, i.realm, usr)

	if err != nil {
		log.Println("Create user error:", err)
		return nil, err
	}

	err = i.client.SetPassword(ctx, token.AccessToken, id, i.realm, password, false)

	if err != nil {
		log.Println("Set password error:", err)
		return nil, err
	}

	keycloakUserID := user.IdPUserID(id)

	return &keycloakUserID, nil
}

func (i *KeycloakIdentityProvider) ConfirmUser(ctx context.Context, id user.IdPUserID) error {
	token, err := i.getToken(ctx)
	if err != nil {
		return err
	}

	usr, err := i.client.GetUserByID(ctx, token.AccessToken, i.realm, id.String())
	if err != nil {
		return err
	}

	if (usr.Enabled != nil && *usr.Enabled) || usr.EmailVerified != nil && *usr.EmailVerified {
		return errors.New("user is already enabled")
	}

	usr.Enabled = gocloak.BoolP(true)
	usr.EmailVerified = gocloak.BoolP(true)

	return i.client.UpdateUser(ctx, token.AccessToken, i.realm, *usr)
}
