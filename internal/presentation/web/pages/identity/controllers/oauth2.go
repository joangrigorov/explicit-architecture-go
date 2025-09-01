package controllers

import (
	"app/internal/presentation/web/services/identity"
	"app/internal/presentation/web/services/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OAuth2 struct {
	authService *identity.AuthenticationService
}

func (a *OAuth2) Callback(c *gin.Context) {
	_, err := a.authService.ObtainToken(session.GetSession(c), c.Query("code"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error":    err,
			"ErrorMsg": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func NewOAuth2(authService *identity.AuthenticationService) *OAuth2 {
	return &OAuth2{authService: authService}
}
