package controllers

import (
	"app/internal/presentation/web/pages/identity/forms"
	"app/internal/presentation/web/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUp struct {
	ids    *services.IdentityService
	client *services.ActivityPlannerClient
}

func NewSignUp(ids *services.IdentityService, client *services.ActivityPlannerClient) *SignUp {
	return &SignUp{ids: ids, client: client}
}

func (u *SignUp) SignUpForm(c *gin.Context) {
	session := services.GetSession(c)
	// TODO have some wrapper around these flashes so I can transport form errors easier
	formErrors := session.Flashes("sign_up_form_errors")
	formError := session.Flashes("sign_up_error")
	values := session.Flashes("sign_up_values")
	if values == nil {
		c.HTML(200, "identity/sign_up", gin.H{
			"Values":     &forms.SignUp{},
			"FormErrors": formErrors,
			"FormError":  formError,
		})
		return
	}

	c.HTML(200, "identity/sign_up", gin.H{
		"Values":     values,
		"FormErrors": formErrors,
		"FormError":  formError,
	})
}

func (u *SignUp) SignUp(c *gin.Context) {
	session := services.GetSession(c)
	req := &forms.SignUp{}
	if err := c.ShouldBind(req); err != nil {
		// Add session flashes
		c.Redirect(http.StatusFound, "/sign-up")
		return
	}
	err := u.client.SignUp(req)
	if err != nil {
		// Add session flashes
		c.Redirect(http.StatusFound, "/sign-up")
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func (u *SignUp) SignInForm(c *gin.Context) {
	c.Redirect(http.StatusFound, u.ids.SignInURL("change-me", "change-me"))
}
