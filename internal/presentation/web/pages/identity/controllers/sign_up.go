package controllers

import (
	"app/internal/presentation/web/pages/identity/forms"
	"app/internal/presentation/web/services"

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
	c.HTML(200, "identity/sign_up", gin.H{
		"Values": &forms.SignUp{},
	})
}

func (u *SignUp) SignUp(c *gin.Context) {
	req := &forms.SignUp{}
	if err := c.ShouldBind(req); err != nil {
		c.HTML(422, "identity/sign_up", gin.H{
			"Error":  err,
			"Values": req,
		})
		return
	}
	err := u.client.SignUp(req)
	if err != nil {
		c.HTML(422, "identity/sign_up", gin.H{
			"Error":  err,
			"Values": req,
		})
		return
	}
	c.Redirect(302, "/")
}

func (u *SignUp) SignInForm(c *gin.Context) {
	c.Redirect(302, u.ids.SignInURL("change-me", "change-me"))
}
