package controllers

import (
	"app/internal/presentation/web/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignOut struct {
	authService *services.AuthenticationService
}

func (o *SignOut) SignOut(c *gin.Context) {
	o.authService.Forget(services.GetSession(c))
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func NewSignOut(authService *services.AuthenticationService) *SignOut {
	return &SignOut{authService: authService}
}
