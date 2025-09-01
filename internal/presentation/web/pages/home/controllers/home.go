package controllers

import (
	"app/internal/presentation/web/services/activity_planner"
	"app/internal/presentation/web/services/identity"
	"app/internal/presentation/web/services/session"

	"github.com/gin-gonic/gin"
)

type Home struct {
	authService *identity.AuthenticationService
	client      *activity_planner.Client
	flash       *session.Flash
}

func NewHome(authService *identity.AuthenticationService, flash *session.Flash, client *activity_planner.Client) *Home {
	return &Home{authService: authService, flash: flash, client: client}
}

func (h *Home) Index(c *gin.Context) {
	sess := session.GetSession(c)

	var profile activity_planner.Profile
	var signedIn bool

	if token, err := h.authService.ActiveSession(sess); err == nil {
		profile, _ = h.client.Me(token)
		signedIn = h.authService.SignedIn(sess)
	}

	c.HTML(200, "home/index", gin.H{
		"Title":    "Welcome Page",
		"Heading":  "Hello from Gin!",
		"Message":  "This page is rendered with Gin and html/template.",
		"SignedIn": signedIn,
		"Alerts":   h.flash.GetAlerts(sess, "landing"),
		"Profile":  profile,
	})
}
