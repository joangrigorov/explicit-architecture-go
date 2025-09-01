package controllers

import (
	"app/internal/presentation/web/pages/identity/forms"
	"app/internal/presentation/web/services/activity_planner"
	"app/internal/presentation/web/services/session"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PasswordSetup struct {
	client *activity_planner.Client
	logger *zap.SugaredLogger
	flash  *session.Flash
}

func (s *PasswordSetup) SetPasswordForm(c *gin.Context) {
	verificationID := c.Query("id")
	token := c.Query("token")
	response, err := s.client.PreflightVerification(verificationID, token)

	sess := session.GetSession(c)
	
	alerts := s.flash.GetAlerts(sess, "password_setup")

	if err != nil {
		c.HTML(http.StatusInternalServerError, "identity/password_setup", gin.H{
			"Alerts":   append(alerts, session.Alert{Kind: session.AlertError, Msg: "Unknown error occurred"}),
			"HideForm": true,
		})
		return
	}

	if response.Error != "" {
		c.HTML(http.StatusInternalServerError, "identity/password_setup", gin.H{
			"Alerts":   append(alerts, session.Alert{Kind: session.AlertError, Msg: response.Error}),
			"HideForm": true,
		})
		return
	}

	if response.Expired {
		c.HTML(http.StatusGone, "identity/password_setup", gin.H{
			"Alerts":   append(alerts, session.Alert{Kind: session.AlertError, Msg: "Password setup session has expired"}),
			"Expired":  true,
			"HideForm": true,
		})
		return
	}

	if !response.ValidCSRF {
		c.HTML(http.StatusBadRequest, "identity/password_setup", gin.H{
			"Alerts":   append(alerts, session.Alert{Kind: session.AlertError, Msg: "CSRF token is invalid"}),
			"HideForm": true,
		})
		return
	}

	formErrors := s.flash.GetFormErrors(sess, forms.PasswordSetup{})

	c.HTML(http.StatusOK, "identity/password_setup", gin.H{
		"MaskedEmail": response.MaskedEmail,
		"id":          verificationID,
		"token":       token,
		"FormErrors":  formErrors,
	})
	return
}

func (s *PasswordSetup) SetPassword(c *gin.Context) {
	form := &forms.PasswordSetup{}
	if err := c.ShouldBind(form); err != nil {
		s.flash.AddFormErrors(session.GetSession(c), *form, err)
		s.flash.AddFormValues(session.GetSession(c), *form)
		log.Println("Params", form.VerificationID, form.Token, err.Error())
		c.Redirect(http.StatusFound, "/set-password?id="+form.VerificationID+"&token="+form.Token)
		return
	}

	if err := s.client.PasswordSetup(form); err != nil {
		s.flash.AddAlert(session.GetSession(c), "password_setup", session.AlertError, err.Error())
		s.flash.AddFormValues(session.GetSession(c), *form)
		c.Redirect(http.StatusFound, "/set-password?id="+form.VerificationID+"&token="+form.Token)
		return
	}

	s.flash.AddAlert(session.GetSession(c), "landing", session.AlertSuccess, "Registration finalized")
	c.Redirect(http.StatusFound, "/")
}

func NewPasswordSetup(
	client *activity_planner.Client,
	logger *zap.SugaredLogger,
	flash *session.Flash,
) *PasswordSetup {
	return &PasswordSetup{client: client, logger: logger, flash: flash}
}
