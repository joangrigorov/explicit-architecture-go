package controllers

import (
	"app/internal/presentation/web/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OAuth2 struct {
	authService *services.AuthenticationService
}

func (a *OAuth2) Callback(c *gin.Context) {
	_, err := a.authService.ObtainToken(services.GetSession(c), c.Query("code"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error":    err,
			"ErrorMsg": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func NewOAuth2(authService *services.AuthenticationService) *OAuth2 {
	return &OAuth2{authService: authService}
}
