package controllers

import (
	"app/internal/presentation/web/services/identity"
	"app/internal/presentation/web/services/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignOut struct {
	authService *identity.AuthenticationService
}

func (o *SignOut) SignOut(c *gin.Context) {
	o.authService.Forget(session.GetSession(c))
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func NewSignOut(authService *identity.AuthenticationService) *SignOut {
	return &SignOut{authService: authService}
}
