package web

import (
	"app/config/web"
	homeC "app/internal/presentation/web/pages/home/controllers"
	id "app/internal/presentation/web/pages/identity/controllers"
	"app/internal/presentation/web/services/session"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
)

func RegisterRoutes(
	r *gin.Engine,
	tmpl *template.Template,
	logger *zap.SugaredLogger,
	store sessions.Store,
	cfg web.Config,

	h *homeC.Home,
	su *id.SignUp,
	so *id.SignOut,
	ps *id.PasswordSetup,
	oauth2 *id.OAuth2,
) {
	r.Use(session.Handler(logger, store, cfg.Session.SessionKey))

	r.StaticFS("/assets", http.Dir("internal/presentation/web/assets"))

	r.SetHTMLTemplate(tmpl)

	r.GET("/", h.Index)
	r.GET("/sign-up", su.SignUpForm)
	r.POST("/sign-up", su.SignUp)
	r.GET("/sign-in", su.SignInForm)
	r.GET("/sign-out", so.SignOut)
	r.GET("/set-password", ps.SetPasswordForm)
	r.POST("/set-password", ps.SetPassword)
	r.GET("/oauth2/callback", oauth2.Callback)
}
