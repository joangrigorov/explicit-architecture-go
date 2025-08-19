package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Context interface {

	// Context returns the request's context
	Context() context.Context

	// JSON for issuing json body responses
	JSON(code int, obj any)

	// AbortWithStatusJSON aborts the handler chain and responds with JSON
	AbortWithStatusJSON(code int, obj interface{})

	// ShouldBindJSON binds the passed struct pointer using json
	ShouldBindJSON(obj any) error

	// ParamString gets a parameter from path as a string
	ParamString(key string) string

	ParamInt(key string) (int, error)

	// NoContent responds with 204
	NoContent()

	// IsJsonRequest to check if the request content type is application/json
	IsJsonRequest() bool

	// IsJsonBodyValid returns true if there is a request body and it is a valid json
	IsJsonBodyValid() bool

	// Next to call the next handler in the chain
	Next()

	IsPost() bool
	IsPut() bool
	IsPatch() bool
	IsGet() bool
	IsDelete() bool
}

type GinContext struct {
	context *gin.Context
}

func (g *GinContext) NoContent() {
	g.context.Status(http.StatusNoContent)
}

func (g *GinContext) Context() context.Context {
	return g.context.Request.Context()
}

func (g *GinContext) JSON(code int, obj any) {
	g.context.JSON(code, obj)
}

func (g *GinContext) AbortWithStatusJSON(code int, obj interface{}) {
	g.context.AbortWithStatusJSON(code, obj)
}

func (g *GinContext) ShouldBindJSON(obj any) error {
	return g.context.ShouldBindJSON(obj)
}

func (g *GinContext) ParamString(key string) string {
	return g.context.Param(key)
}

func (g *GinContext) ParamInt(key string) (int, error) {
	raw := g.context.Param("id")
	param, err := strconv.Atoi(raw)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("param error: %s: %s cannot be converted to integer", key, raw))
	}
	return param, err
}

func (g *GinContext) IsPost() bool {
	return g.context.Request.Method == http.MethodPost
}

func (g *GinContext) IsPut() bool {
	return g.context.Request.Method == http.MethodPut
}

func (g *GinContext) IsPatch() bool {
	return g.context.Request.Method == http.MethodPatch
}

func (g *GinContext) IsGet() bool {
	return g.context.Request.Method == http.MethodGet
}

func (g *GinContext) IsDelete() bool {
	return g.context.Request.Method == http.MethodDelete
}

func (g *GinContext) IsJsonRequest() bool {
	return strings.HasPrefix(g.context.GetHeader("Content-Type"), "application/json")
}

func (g *GinContext) IsJsonBodyValid() bool {
	if !g.IsJsonRequest() {
		return false
	}

	bodyBytes, err := io.ReadAll(g.context.Request.Body)
	if err != nil {
		return false
	}

	// Restore body for downstream handlers
	g.context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Check JSON validity
	var js json.RawMessage
	if err := json.Unmarshal(bodyBytes, &js); err != nil {
		return false
	}

	return true
}

func (g *GinContext) Next() {
	g.context.Next()
}

func NewGinContext(c *gin.Context) *GinContext {
	return &GinContext{context: c}
}
