package responses

import (
	"app/internal/infrastructure/framework/validation"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type fakeFieldError struct {
	field           string
	structNamespace string
	tag             string
	param           string
	actualTag       string
}

func (f fakeFieldError) Tag() string             { return f.tag }
func (f fakeFieldError) ActualTag() string       { return f.actualTag }
func (f fakeFieldError) Param() string           { return f.param }
func (f fakeFieldError) Field() string           { return f.field }
func (f fakeFieldError) StructField() string     { return f.field }
func (f fakeFieldError) StructNamespace() string { return f.structNamespace }
func (f fakeFieldError) Namespace() string       { return f.structNamespace }
func (f fakeFieldError) Kind() reflect.Kind      { return reflect.String }
func (f fakeFieldError) Type() reflect.Type      { return reflect.TypeOf("") }
func (f fakeFieldError) Value() interface{}      { return nil }
func (f fakeFieldError) ParamInt() int64         { return 0 }
func (f fakeFieldError) Translate(trans ut.Translator) string {
	return "Invalid (fake message)"
}
func (f fakeFieldError) Error() string { return f.Translate(nil) }

func TestRender(t *testing.T) {
	s := ""

	t.Run("render validation errors", func(t *testing.T) {
		type subExample struct {
			Field1 *string `json:"sub_field1"`
			Hidden string  `json:"-"`
		}

		type example struct {
			Field1 *string     `json:"field1"`
			Field2 string      `json:"field2"`
			Field3 *subExample `json:"field3"`
		}

		ve := validator.ValidationErrors{
			fakeFieldError{
				field:           "Field1",
				structNamespace: "example.Field1",
				tag:             "required",
			},
			fakeFieldError{
				field:           "Field1",
				structNamespace: "example.Field3.Field1",
				tag:             "required",
			},
			fakeFieldError{
				field:           "Hidden",
				structNamespace: "example.Field3.Hidden",
				tag:             "required",
			},
		}

		res := Render(validation.NewTranslator(), ve, &example{
			Field1: &s,
			Field2: "valid",
			Field3: &subExample{
				Field1: &s,
			},
		})

		assert.Equal(t, ErrorsResponse{
			Errors: []fieldError{
				{
					Field:   "field1",
					Rule:    "required",
					Message: "Invalid (fake message)",
				},
				{
					Field:   "field3.sub_field1",
					Rule:    "required",
					Message: "Invalid (fake message)",
				},
				{
					Field:   "field3.Hidden",
					Rule:    "required",
					Message: "Invalid (fake message)",
				},
			},
		}, res)
	})

	t.Run("panic on root mismatch", func(t *testing.T) {
		type wrong struct {
			Field1 *string `json:"field1"`
		}

		type example struct {
			Field1 *string `json:"field1"`
		}

		ve := validator.ValidationErrors{
			fakeFieldError{
				field:           "Field1",
				structNamespace: "example.Field1",
				tag:             "required",
			},
		}
		assert.Panics(t, func() {
			Render(validation.NewTranslator(), ve, &wrong{
				Field1: &s,
			})
		})
	})

	t.Run("panic on non-existing field", func(t *testing.T) {
		type example struct {
			Field *string `json:"field"`
		}

		ve := validator.ValidationErrors{
			fakeFieldError{
				field:           "Invalid",
				structNamespace: "example.Invalid",
				tag:             "required",
			},
		}
		assert.Panics(t, func() {
			Render(validation.NewTranslator(), ve, &example{
				Field: &s,
			})
		})
	})

	t.Run("panic on non-existing field", func(t *testing.T) {
		type example struct {
			Field *string `json:"field"`
		}

		ve := validator.ValidationErrors{
			fakeFieldError{
				field:           "Invalid",
				structNamespace: "example.Invalid",
				tag:             "required",
			},
		}
		assert.Panics(t, func() {
			Render(validation.NewTranslator(), ve, &example{
				Field: &s,
			})
		})
	})

	t.Run("panic on impossible nesting", func(t *testing.T) {
		type example struct {
			Field1 *string `json:"field1"`
		}

		ve := validator.ValidationErrors{
			fakeFieldError{
				field:           "Field1",
				structNamespace: "example.Field1.TooDeep",
				tag:             "required",
			},
		}

		assert.Panics(t, func() {
			Render(validation.NewTranslator(), ve, &example{
				Field1: &s,
			})
		})
	})

	t.Run("render json type errors", func(t *testing.T) {
		err := json.UnmarshalTypeError{Field: "some_field"}
		assert.Equal(t, ErrorsResponse{Errors: []fieldError{
			{
				Field:   "some_field",
				Rule:    "type_error",
				Message: "Type error for field some_field",
			},
		}}, Render(validation.NewTranslator(), &err, nil))
	})

	t.Run("single error", func(t *testing.T) {
		errMsg := "some error"
		err := errors.New(errMsg)
		assert.Equal(t, ErrorsResponse{Error: &errMsg}, Render(validation.NewTranslator(), err, nil))
	})
}
