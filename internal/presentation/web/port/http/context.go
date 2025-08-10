package http

import (
	"context"
)

type Context interface {

	// Context returns the request's context
	Context() context.Context

	// JSON for issuing json body responses
	JSON(code int, obj any)

	// ShouldBindJSON binds the passed struct pointer using json
	ShouldBindJSON(obj any) error

	// Param gets a parameter from path as a string
	Param(key string) string

	// NoContent responds with 204
	NoContent()
}

type Handler func(Context)
