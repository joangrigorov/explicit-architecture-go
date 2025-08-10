package rules

import (
	ut "app/mock/ext/go-playground/universal-translator"
	"app/mock/ext/go-playground/validator"
	"github.com/aarondl/null/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestNewNotNull(t *testing.T) {
	assert.NotNil(t, NewNotNull())
}

func TestNotNull_Tag(t *testing.T) {
	assert.Equal(t, "notnull", NewNotNull().Tag())
}

func TestNotNull_Translate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rule := NewNotNull()

	translator := ut.NewMockTranslator(ctrl)
	fieldError := validator.NewMockFieldError(ctrl)

	fieldError.
		EXPECT().
		Field().
		Return("a_field")

	translator.
		EXPECT().
		T("notnull", "a_field").
		Return("irrelevant", nil)

	rule.Translate(translator, fieldError)
}

func TestNotNull_RegisterTranslations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rule := NewNotNull()
	translator := ut.NewMockTranslator(ctrl)

	translator.
		EXPECT().
		Add("notnull", "The {0} field must not be null", true)

	assert.Nil(t, rule.RegisterTranslations(translator))
}

func TestNotNull_Validate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rule := NewNotNull()

	t.Run("validates a valid not null", func(t *testing.T) {
		value := null.NewString("a value", true)

		fl := validator.NewMockFieldLevel(ctrl)
		fl.EXPECT().Field().Return(reflect.ValueOf(value))

		assert.True(t, rule.Validate(fl))
	})

	t.Run("validates a value that is not set", func(t *testing.T) {
		value := null.NewString("a value", true)
		value.Set = false

		fl := validator.NewMockFieldLevel(ctrl)
		fl.EXPECT().Field().Return(reflect.ValueOf(value))

		assert.True(t, rule.Validate(fl))
	})

	t.Run("validation fails on type mismatch", func(t *testing.T) {
		value := "we require a type from the null package"

		fl := validator.NewMockFieldLevel(ctrl)
		fl.EXPECT().Field().Return(reflect.ValueOf(value))

		assert.False(t, rule.Validate(fl))
	})

	t.Run("validation fails if value is set as invalid", func(t *testing.T) {
		value := null.NewString("a value", false)

		fl := validator.NewMockFieldLevel(ctrl)
		fl.EXPECT().Field().Return(reflect.ValueOf(value))

		assert.False(t, rule.Validate(fl))
	})
}
