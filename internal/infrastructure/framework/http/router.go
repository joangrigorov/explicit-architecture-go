package http

import (
	ctx "app/internal/presentation/web/port/http"
	"github.com/gin-gonic/gin"
)

func wrapHandler(h ctx.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(NewGinContext(c))
	}
}

type group struct {
	rg *gin.RouterGroup
}

func (r *group) Group(path string) ctx.RouterGroup {
	ginRg := r.rg.Group(path)
	return &group{rg: ginRg}
}

func (r *group) POST(path string, handler ctx.Handler) {
	r.rg.POST(path, wrapHandler(handler))
}

func (r *group) GET(path string, handler ctx.Handler) {
	r.rg.GET(path, wrapHandler(handler))
}

func (r *group) DELETE(path string, handler ctx.Handler) {
	r.rg.DELETE(path, wrapHandler(handler))
}

func (r *group) PATCH(path string, handler ctx.Handler) {
	r.rg.PATCH(path, wrapHandler(handler))
}

func (r *group) PUT(path string, handler ctx.Handler) {
	r.rg.PUT(path, wrapHandler(handler))
}

type Router struct {
	engine *gin.Engine
}

func (r *Router) Group(path string) ctx.RouterGroup {
	ginRg := r.engine.Group(path)
	return &group{rg: ginRg}
}

func (r *Router) POST(path string, handler ctx.Handler) {
	r.engine.POST(path, wrapHandler(handler))
}

func (r *Router) GET(path string, handler ctx.Handler) {
	r.engine.GET(path, wrapHandler(handler))
}

func (r *Router) DELETE(path string, handler ctx.Handler) {
	r.engine.DELETE(path, wrapHandler(handler))
}

func (r *Router) PATCH(path string, handler ctx.Handler) {
	r.engine.PATCH(path, wrapHandler(handler))
}

func (r *Router) PUT(path string, handler ctx.Handler) {
	r.engine.PUT(path, wrapHandler(handler))
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func NewRouter(engine *gin.Engine) ctx.Router {
	return &Router{
		engine: engine,
	}
}
