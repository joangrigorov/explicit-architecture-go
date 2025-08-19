package http

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	// POST method handling on a relative path
	POST(string, ...Handler)

	// GET method handling on a relative path
	GET(string, ...Handler)

	// DELETE method handling on a relative path
	DELETE(string, ...Handler)

	// PATCH method handling on a relative path
	PATCH(string, ...Handler)

	// PUT method handling on a relative path
	PUT(string, ...Handler)

	// Group routes under a common relative path
	Group(string) Router

	// Use global handler
	Use(...Handler)
}

type Handler func(Context)

func adaptHandlers(handlers []Handler) []gin.HandlerFunc {
	var wrapped []gin.HandlerFunc

	for _, handler := range handlers {
		wrapped = append(wrapped, func(c *gin.Context) {
			handler(NewGinContext(c))
		})
	}
	return wrapped
}

type DefaultRouter struct {
	router gin.IRouter
}

func (r *DefaultRouter) Group(path string) Router {
	return &DefaultRouter{router: r.router.Group(path)}
}

func (r *DefaultRouter) Use(handler ...Handler) {
	r.router.Use(adaptHandlers(handler)...)
}

func (r *DefaultRouter) POST(path string, handlers ...Handler) {
	r.router.POST(path, adaptHandlers(handlers)...)
}

func (r *DefaultRouter) GET(path string, handlers ...Handler) {
	r.router.GET(path, adaptHandlers(handlers)...)
}

func (r *DefaultRouter) DELETE(path string, handlers ...Handler) {
	r.router.DELETE(path, adaptHandlers(handlers)...)
}

func (r *DefaultRouter) PATCH(path string, handlers ...Handler) {
	r.router.PATCH(path, adaptHandlers(handlers)...)
}

func (r *DefaultRouter) PUT(path string, handlers ...Handler) {
	r.router.PUT(path, adaptHandlers(handlers)...)
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func NewRouter(engine *gin.Engine) Router {
	return &DefaultRouter{
		router: engine,
	}
}
