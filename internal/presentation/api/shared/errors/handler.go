package errors

import (
	"app/config/api"
	"app/internal/core/port/errors"
	ctx "app/internal/infrastructure/framework/http"
	"app/internal/infrastructure/framework/support"
	"encoding/json"
	stdErrs "errors"
	"fmt"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	tr    ut.Translator
	debug bool
}

type fieldError struct {
	Field   string `json:"field"`
	Rule    string `json:"rule"`
	Message string `json:"message"`
}

type response struct {
	Error  string       `json:"error,omitempty"`
	Stack  []string     `json:"stack,omitempty"`
	Errors []fieldError `json:"errors,omitempty"`
}

func (h *Handler) RenderApplicationError(c ctx.Context, err error) {
	var appErr errors.Error
	if stdErrs.As(err, &appErr) {
		var stack []string
		if h.debug {
			stack = appErr.PrettyPrint()
		}
		c.JSON(httpStatusCode(appErr), response{Error: appErr.Error(), Stack: stack})
		return
	}

	c.JSON(http.StatusInternalServerError, response{Error: "Unknown error occurred"})
}

func (h *Handler) RenderRequestError(c ctx.Context, err error, request interface{}) {
	var ve validator.ValidationErrors
	if stdErrs.As(err, &ve) {
		c.JSON(http.StatusUnprocessableEntity, h.renderValidationErrors(ve, request))
		return
	}

	var ute *json.UnmarshalTypeError
	if stdErrs.As(err, &ute) {
		c.JSON(http.StatusUnprocessableEntity, h.renderTypeErrors(ute))
		return
	}

	c.JSON(http.StatusInternalServerError, "Unknown error occurred")
	return
}

func httpStatusCode(err errors.Error) int {
	switch err.Code() {
	case errors.ErrValidation:
		return http.StatusUnprocessableEntity
	case errors.ErrNotFound:
		return http.StatusNotFound
	case errors.ErrConflict:
		return http.StatusConflict
	case errors.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func (h *Handler) renderTypeErrors(ute *json.UnmarshalTypeError) response {
	return response{
		Errors: []fieldError{
			{
				Field:   ute.Field,
				Rule:    "type_error",
				Message: fmt.Sprintf("Type error for field %s", ute.Field),
			},
		},
	}
}

func (h *Handler) renderValidationErrors(ve validator.ValidationErrors, req interface{}) response {
	fieldErrors := make([]fieldError, len(ve))
	for i, fe := range ve {
		fieldErrors[i] = fieldError{
			Field:   support.TagFieldName(fe, req, "json"),
			Message: fe.Translate(h.tr),
			Rule:    fe.Tag(),
		}
	}
	return response{
		Errors: fieldErrors,
	}
}

func NewHandler(tr ut.Translator, cfg api.Config) *Handler {
	return &Handler{tr: tr, debug: cfg.App.Debug}
}
