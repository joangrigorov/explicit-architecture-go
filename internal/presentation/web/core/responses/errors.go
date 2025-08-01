package responses

import (
	"app/internal/infrastructure/framework/validation"
	ctx "app/internal/presentation/web/port/http"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type defaultError struct {
	Error string `json:"string"`
}

func newDefaultError(err error) defaultError {
	return defaultError{
		Error: err.Error(),
	}
}

func UnprocessableEntity(c ctx.Context, err error, req interface{}) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		c.JSON(http.StatusUnprocessableEntity, validation.Render(ve, req))
		return
	}
	c.JSON(http.StatusUnprocessableEntity, newDefaultError(err))
}

func InternalServerError(c ctx.Context, err error) {
	c.JSON(http.StatusInternalServerError, newDefaultError(err))
}

func BadRequest(c ctx.Context, err error) {
	c.JSON(http.StatusBadRequest, newDefaultError(err))
}

func NotFound(c ctx.Context, err error) {
	c.JSON(http.StatusNotFound, newDefaultError(err))
}
