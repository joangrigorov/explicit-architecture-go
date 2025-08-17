package responses

import (
	ctx "app/internal/presentation/api/port/http"
	"net/http"
)

type DefaultError struct {
	Error string `json:"error"`
}

func NewDefaultError(err error) DefaultError {
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
