package controllers

import (
	"app/internal/presentation/web/pages/identity/forms"
	"app/internal/presentation/web/services/activity_planner"
	"app/internal/presentation/web/services/identity"
	"app/internal/presentation/web/services/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUp struct {
	idp    *identity.IdPService
	flash  *session.Flash
	client *activity_planner.Client
}

func NewSignUp(idp *identity.IdPService, flash *session.Flash, client *activity_planner.Client) *SignUp {
	return &SignUp{idp: idp, flash: flash, client: client}
}

func (u *SignUp) SignUpForm(c *gin.Context) {
	ses := session.GetSession(c)
	c.HTML(200, "identity/sign_up", gin.H{
		"Values":     u.flash.GetFormValues(ses, forms.SignUp{}),
		"FormErrors": u.flash.GetFormErrors(ses, forms.SignUp{}),
		"Alerts":     u.flash.GetAlerts(ses, "sign_up"),
	})
}

func (u *SignUp) SignUp(c *gin.Context) {
	form := &forms.SignUp{}
	if err := c.ShouldBind(form); err != nil {
		u.flash.AddFormErrors(session.GetSession(c), *form, err)
		u.flash.AddFormValues(session.GetSession(c), *form)
		c.Redirect(http.StatusFound, "/sign-up")
		return
	}
	err := u.client.SignUp(form)
	if err != nil {
		u.flash.AddAlert(session.GetSession(c), "sign_up", session.AlertError, err.Error())
		u.flash.AddFormValues(session.GetSession(c), *form)
		c.Redirect(http.StatusFound, "/sign-up")
		return
	}
	u.flash.AddAlert(
		session.GetSession(c),
		"landing",
		session.AlertSuccess,
		"Sign-up successful. Please check your email for password setup email!",
	)
	c.Redirect(http.StatusFound, "/")
}

func (u *SignUp) SignInForm(c *gin.Context) {
	c.Redirect(http.StatusFound, u.idp.SignInURL("change-me", "change-me"))
}
