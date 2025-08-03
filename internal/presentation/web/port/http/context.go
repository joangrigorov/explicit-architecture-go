package http

import (
	"context"
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
