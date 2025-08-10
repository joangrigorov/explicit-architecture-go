package http

import (
	ctx "app/internal/presentation/web/port/http"

	"github.com/gin-gonic/gin"
)

func adaptHandlers(handlers []ctx.Handler) []gin.HandlerFunc {
	var wrapped []gin.HandlerFunc

	for _, handler := range handlers {
		wrapped = append(wrapped, func(c *gin.Context) {
			handler(NewGinContext(c))
		})
	}
	return wrapped
}

type Router struct {
	router gin.IRouter
}

func (r *Router) Group(path string) ctx.Router {
	return &Router{router: r.router.Group(path)}
}

func (r *Router) Use(handler ...ctx.Handler) {
	r.router.Use(adaptHandlers(handler)...)
}

func (r *Router) POST(path string, handlers ...ctx.Handler) {
	r.router.POST(path, adaptHandlers(handlers)...)
}

func (r *Router) GET(path string, handlers ...ctx.Handler) {
	r.router.GET(path, adaptHandlers(handlers)...)
}

func (r *Router) DELETE(path string, handlers ...ctx.Handler) {
	r.router.DELETE(path, adaptHandlers(handlers)...)
}

func (r *Router) PATCH(path string, handlers ...ctx.Handler) {
	r.router.PATCH(path, adaptHandlers(handlers)...)
}

func (r *Router) PUT(path string, handlers ...ctx.Handler) {
	r.router.PUT(path, adaptHandlers(handlers)...)
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func NewRouter(engine *gin.Engine) ctx.Router {
	return &Router{
		router: engine,
	}
}
