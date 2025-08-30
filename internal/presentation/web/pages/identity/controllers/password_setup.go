package controllers

import (
	"app/internal/presentation/web/pages/identity/forms"
	"app/internal/presentation/web/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PasswordSetup struct {
	client *services.ActivityPlannerClient
	logger *zap.SugaredLogger
}

func (s *PasswordSetup) SetPasswordForm(c *gin.Context) {
	verificationID := c.Query("id")
	token := c.Query("token")
	response, err := s.client.PreflightVerification(verificationID, token)

	if err != nil {
		s.logger.Error("error while verifying token", zap.Error(err))
		c.HTML(http.StatusInternalServerError, "identity/password_setup", gin.H{
			"Error": "Unknown error occurred",
		})
		return
	}

	if response.Error != "" {
		c.HTML(http.StatusInternalServerError, "identity/password_setup", gin.H{
			"Error":   response.Error,
			"Expired": false,
		})
		return
	}

	if response.Expired {
		c.HTML(http.StatusGone, "identity/password_setup", gin.H{
			"Error":   "Token is expired",
			"Expired": true,
		})
		return
	}

	if !response.ValidCSRF {
		c.HTML(http.StatusBadRequest, "identity/password_setup", gin.H{
			"Error":   "CSRF token is invalid",
			"Expired": false,
		})
		return
	}

	c.HTML(http.StatusOK, "identity/password_setup", gin.H{
		"MaskedEmail": response.MaskedEmail,
		"ID":          verificationID,
		"Token":       token,
	})
	return
}

func (s *PasswordSetup) SetPassword(c *gin.Context) {
	req := &forms.PasswordSetup{}
	if err := c.ShouldBind(req); err != nil {
		// TODO render form errors properly
		c.HTML(422, "identity/password_setup", gin.H{
			"Error": err.Error(),
		})
		return
	}

	if err := s.client.PasswordSetup(req); err != nil {
		// TODO save error in session and redirect to the form
		c.HTML(422, "identity/password_setup", gin.H{
			"Error": err.Error(),
		})
		return
	}

	// TODO save a success state in session and display it in the sign in form
	c.Redirect(http.StatusFound, "/sign-in")
}

func NewPasswordSetup(client *services.ActivityPlannerClient, logger *zap.SugaredLogger) *PasswordSetup {
	return &PasswordSetup{client: client, logger: logger}
}
