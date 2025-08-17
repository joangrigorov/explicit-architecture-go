package idp

import (
	"app/config"
	"app/internal/core/port/idp"
	"app/internal/core/shared_kernel/domain"
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

func NewKeycloakIdentityProvider(client *gocloak.GoCloak, cfg *config.Config) idp.IdentityProvider {
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
	userID domain.UserID,
	username string,
	email string,
	password string,
) (*domain.IdPUserId, error) {
	token, err := i.getToken(ctx)
	if err != nil {
		return nil, err
	}

	user := gocloak.User{
		Email:    gocloak.StringP(email),
		Username: gocloak.StringP(username),
		Enabled:  gocloak.BoolP(false),
		Attributes: &map[string][]string{
			"app_user_id": {userID.String()},
		},
	}

	id, err := i.client.CreateUser(ctx, token.AccessToken, i.realm, user)

	if err != nil {
		log.Println("Create user error:", err)
		return nil, err
	}

	err = i.client.SetPassword(ctx, token.AccessToken, id, i.realm, password, false)

	if err != nil {
		log.Println("Set password error:", err)
		return nil, err
	}

	keycloakUserID := domain.IdPUserId(id)

	return &keycloakUserID, nil
}

func (i *KeycloakIdentityProvider) ConfirmUser(ctx context.Context, id domain.IdPUserId) error {
	token, err := i.getToken(ctx)
	if err != nil {
		return err
	}

	user, err := i.client.GetUserByID(ctx, token.AccessToken, i.realm, id.String())
	if err != nil {
		return err
	}

	if (user.Enabled != nil && *user.Enabled) || user.EmailVerified != nil && *user.EmailVerified {
		return errors.New("user is already enabled")
	}

	user.Enabled = gocloak.BoolP(true)
	user.EmailVerified = gocloak.BoolP(true)

	return i.client.UpdateUser(ctx, token.AccessToken, i.realm, *user)
}
