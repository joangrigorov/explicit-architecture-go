package web

import (
	homeC "app/internal/presentation/web/pages/home/controllers"
	"html/template"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	h *homeC.Home,
) {
	tmpl := template.Must(template.New("").ParseGlob("internal/presentation/web/layout/*.gohtml"))
	template.Must(tmpl.ParseGlob("internal/presentation/web/pages/**/**/*.gohtml"))

	r.SetHTMLTemplate(tmpl)

	r.GET("/", h.Index)
}
