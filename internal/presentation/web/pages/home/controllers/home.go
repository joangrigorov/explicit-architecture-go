package controllers

import (
	"app/internal/presentation/web/services"

	"github.com/gin-gonic/gin"
)

type Home struct {
	authService *services.AuthenticationService
}

func NewHome(authService *services.AuthenticationService) *Home {
	return &Home{authService: authService}
}

func (h *Home) Index(c *gin.Context) {
	c.HTML(200, "home/index", gin.H{
		"Title":    "Welcome Page",
		"Heading":  "Hello from Gin!",
		"Message":  "This page is rendered with Gin and html/template.",
		"SignedIn": h.authService.SignedIn(services.GetSession(c)),
	})
}
