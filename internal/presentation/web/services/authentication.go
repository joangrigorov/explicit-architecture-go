package services

import (
	"encoding/gob"
	"log"
	"time"

	"github.com/gorilla/sessions"
)

const authSessionKey = "auth"

func init() {
	gob.Register(AuthenticationSession{})
}

type AuthenticationService struct {
	ids *IdentityService
}

func (s *AuthenticationService) ObtainToken(session *sessions.Session, code string) (AuthenticationSession, error) {
	token, err := s.ids.ExchangeCodeForToken(code)
	if err != nil {
		return AuthenticationSession{}, err
	}

	authSession := AuthenticationSession{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
	}

	session.Values["auth"] = authSession

	return authSession, nil
}

func (s *AuthenticationService) ActiveSession(session *sessions.Session) (AuthenticationSession, error) {
	authSession, ok := session.Values[authSessionKey].(AuthenticationSession)
	if !ok {
		token, err := s.ids.RefreshToken(authSession.RefreshToken)
		if err != nil {
			delete(session.Values, authSessionKey)
			return AuthenticationSession{}, err
		}
		authenticationSession := AuthenticationSession{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			ExpiresAt:    time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
		}
		session.Values[authSessionKey] = authenticationSession
		return authenticationSession, nil
	}

	return authSession, nil
}

func (s *AuthenticationService) SignedIn(session *sessions.Session) bool {
	_, err := s.ActiveSession(session)
	return err == nil
}

func (s *AuthenticationService) Forget(session *sessions.Session) {
	activeSession, err := s.ActiveSession(session)
	if err == nil {
		if err = s.ids.Logout(activeSession.RefreshToken); err != nil {
			log.Println("logout failed", err)
		}
	}
	delete(session.Values, authSessionKey)
}

type AuthenticationSession struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (s AuthenticationSession) Expired() bool {
	return time.Now().After(s.ExpiresAt)
}

func NewAuthenticationService(ids *IdentityService) *AuthenticationService {
	return &AuthenticationService{ids: ids}
}
