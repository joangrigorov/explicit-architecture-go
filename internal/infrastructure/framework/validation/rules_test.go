package validation

import (
	utMock "app/mock/ext/go-playground/universal-translator"
	"app/mock/infrastructure/framework/validation"
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRegisterRules(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		rule := validation.NewMockRule(ctrl)
		rule.EXPECT().
			Tag().Return("example")
		rule.EXPECT().
			RegisterTranslations(gomock.Any()).
			Return(nil).
			AnyTimes()

		customRules = []Rule{rule}

		v := validator.New()
		tr := utMock.NewMockTranslator(ctrl)
		tr.EXPECT().
			Add(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		tr.EXPECT().
			AddCardinal(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()

		require.NotPanics(t, func() {
			RegisterRules(tr, v)
		})
	})

	t.Run("empty tag causes panic", func(t *testing.T) {
		rule := validation.NewMockRule(ctrl)
		rule.EXPECT().
			Tag().Return("")

		customRules = []Rule{rule}

		v := validator.New()
		tr := utMock.NewMockTranslator(ctrl)

		require.Panics(t, func() {
			RegisterRules(tr, v)
		})
	})

	t.Run("nil translator causes panic", func(t *testing.T) {
		rule := validation.NewMockRule(ctrl)
		rule.EXPECT().
			Tag().Return("example")
		rule.EXPECT().
			RegisterTranslations(gomock.Any()).
			Return(nil).
			AnyTimes()

		// Override global customRules
		customRules = []Rule{rule}

		require.Panics(t, func() {
			RegisterRules(nil, nil)
		})
	})
}

func TestNewTranslator(t *testing.T) {
	assert.NotNil(t, NewTranslator())
}

func TestNewValidatorValidate(t *testing.T) {
	assert.NotNil(t, NewValidatorValidate())
}

type dummyValidator struct{}

func (d dummyValidator) ValidateStruct(any) error { return nil }
func (d dummyValidator) Engine() any              { return struct{}{} }

func TestNewValidatorValidate_Panics(t *testing.T) {
	original := binding.Validator
	binding.Validator = dummyValidator{}

	defer func() { binding.Validator = original }()

	require.PanicsWithValue(t, "Cannot register validator", func() {
		NewValidatorValidate()
	})
}
