package http

import (
	"context"
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	context *gin.Context
}

func (g *GinContext) NoContent(code int) {
	g.context.Status(code)
}

func (g *GinContext) Context() context.Context {
	return g.context.Request.Context()
}

func (g *GinContext) JSON(code int, obj any) {
	g.context.JSON(code, obj)
}

func (g *GinContext) ShouldBindJSON(obj any) error {
	return g.context.ShouldBindJSON(obj)
}

func (g *GinContext) Param(key string) string {
	return g.context.Param(key)
}

func NewGinContext(c *gin.Context) *GinContext {
	return &GinContext{context: c}
}
