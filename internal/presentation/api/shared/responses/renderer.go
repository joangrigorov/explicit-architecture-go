package responses

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

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
			Field:   getJSONFieldName(fe, req),
			Message: fe.Translate(tr),
			Rule:    fe.Tag(),
		}
	}
	return ErrorsResponse{
		Errors: fieldErrors,
	}
}

// getJSONFieldName returns the JSON tag-based field path for the given validator.FieldError
func getJSONFieldName(fe validator.FieldError, root interface{}) string {
	ns := fe.StructNamespace()
	parts := strings.Split(ns, ".")

	t := reflect.TypeOf(root)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	rootName := t.Name()
	if len(parts) == 0 || parts[0] != rootName {
		panic(fmt.Sprintf("validation error namespace root %q does not match provided type %q", parts[0], rootName))
	}

	if len(parts) > 0 {
		parts = parts[1:]
	}

	var jsonPath []string

	for _, part := range parts {
		if t.Kind() != reflect.Struct {
			panic("nesting mismatch")
		}

		var field reflect.StructField
		found := false

		for i := 0; i < t.NumField(); i++ {
			field = t.Field(i)
			if field.Name == part {
				found = true
				break
			}
		}

		if !found {
			panic(fmt.Sprintf("field %q not found in struct %s", part, t.Name()))
		}

		jsonTag := field.Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName == "-" || jsonName == "" {
			jsonName = field.Name
		}

		jsonPath = append(jsonPath, jsonName)
		t = field.Type
		for t.Kind() == reflect.Pointer {
			t = t.Elem()
		}
	}

	return strings.Join(jsonPath, ".")
}
