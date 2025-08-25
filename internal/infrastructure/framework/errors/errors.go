package errors

import (
	"app/internal/core/port/errors"
	"runtime"
)

type ErrorFactory struct{}

func NewErrorFactory() errors.ErrorFactory {
	return &ErrorFactory{}
}

func (e *ErrorFactory) New(code errors.Code, message string, prev error) errors.Error {
	const depth = 32
	pcs := make([]uintptr, depth)
	n := runtime.Callers(3, pcs)
	return Error{
		code:    code,
		message: message,
		prev:    prev,
		stack:   pcs[:n],
	}
}

type Error struct {
	code    errors.Code
	message string
	prev    error
	stack   []uintptr
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() errors.Code {
	return e.code
}

func (e Error) Message() string {
	return e.message
}

func (e Error) Previous() error {
	return e.prev
}

func (e Error) StackTrace() []uintptr {
	return e.stack
}
