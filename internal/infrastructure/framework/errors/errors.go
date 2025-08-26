package errors

import (
	"app/internal/core/port/errors"
	"fmt"
	"runtime"
)

type ErrorFactory struct{}

func NewErrorFactory() errors.ErrorFactory {
	return &ErrorFactory{}
}

func (e *ErrorFactory) New(code errors.Code, message string, prev error) errors.Error {
	const depth = 32
	pcs := make(errors.StackTrace, depth)
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
	stack   errors.StackTrace
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

func (e Error) StackTrace() errors.StackTrace {
	return e.stack
}

func (e Error) PrettyPrint() []string {
	frames := runtime.CallersFrames(e.stack)
	var trace []string
	for {
		f, more := frames.Next()
		trace = append(trace, fmt.Sprintf("%s\n\t%s:%d", f.Function, f.File, f.Line))
		if !more {
			break
		}
	}
	return trace
}
