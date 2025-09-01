package responses

import (
	. "app/internal/core/port/errors"
	ctx "app/internal/infrastructure/framework/http"
	"errors"
	"net/http"
)

type ErrorHandler struct {
}

// TODO convert this into a factory and make a function that automatically
// TODO detects if the error is a controlled application error,
// TODO a validation error or generic error of any other kind.
type DefaultError struct {
	Error string   `json:"error"`
	Stack []string `json:"stack,omitempty"`
}

func NewDefaultError(err error) DefaultError {
	var appErr Error
	if errors.As(err, &appErr) {
		// TODO only add the stack trace if app is set in debug mode
		return DefaultError{appErr.Error(), appErr.PrettyPrint()}
	}

	return DefaultError{
		Error: err.Error(),
	}
}

func UnprocessableEntity(c ctx.Context, res interface{}) {
	c.JSON(http.StatusUnprocessableEntity, res)
}

func InternalServerError(c ctx.Context, res interface{}) {
	c.JSON(http.StatusInternalServerError, res)
}

func BadRequest(c ctx.Context, res interface{}) {
	c.JSON(http.StatusBadRequest, res)
}

func NotFound(c ctx.Context, res interface{}) {
	c.JSON(http.StatusNotFound, res)
}

func Conflict(c ctx.Context, res interface{}) {
	c.JSON(http.StatusConflict, res)
}

func Gone(c ctx.Context, res interface{}) {
	c.JSON(http.StatusGone, res)
}
