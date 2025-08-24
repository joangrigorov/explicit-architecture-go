package web

import (
	homeC "app/internal/presentation/web/pages/home/controllers"
	id "app/internal/presentation/web/pages/identity/controllers"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	tmpl *template.Template,

	h *homeC.Home,
	su *id.SignUp,
) {
	r.StaticFS("/assets", http.Dir("internal/presentation/web/assets"))

	r.SetHTMLTemplate(tmpl)

	r.GET("/", h.Index)
	r.GET("/sign-up", su.SignUpForm)
	r.POST("/sign-up", su.SignUp)
	r.GET("/sign-in", su.SignInForm)
}
