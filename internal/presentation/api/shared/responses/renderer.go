package responses

import (
	"app/internal/infrastructure/framework/support"
	"encoding/json"
	"errors"
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type fieldError struct {
	Field   string `json:"field"`
	Rule    string `json:"rule"`
	Message string `json:"message"`
}

type ErrorsResponse struct {
	Errors []fieldError `json:"errors"`
	Error  *string      `json:"error,omitempty"`
}

// Deprecated use errors.Handler instead
func Render(tr ut.Translator, err error, req interface{}) ErrorsResponse {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return renderValidationErrors(tr, ve, req)
	}

	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		return renderTypeErrors(ute)
	}

	errorAsString := err.Error()
	return ErrorsResponse{Error: &errorAsString}
}

func renderTypeErrors(ute *json.UnmarshalTypeError) ErrorsResponse {
	return ErrorsResponse{
		Errors: []fieldError{
			{
				Field:   ute.Field,
				Rule:    "type_error",
				Message: fmt.Sprintf("Type error for field %s", ute.Field),
			},
		},
	}
}

func renderValidationErrors(tr ut.Translator, ve validator.ValidationErrors, req interface{}) ErrorsResponse {
	fieldErrors := make([]fieldError, len(ve))
	for i, fe := range ve {
		fieldErrors[i] = fieldError{
			Field:   support.TagFieldName(fe, req, "json"),
			Message: fe.Translate(tr),
			Rule:    fe.Tag(),
		}
	}
	return ErrorsResponse{
		Errors: fieldErrors,
	}
}
