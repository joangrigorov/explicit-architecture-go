package post

import (
	"app/mock/core/component/blog/application/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestNewController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := repositories.NewMockPostRepository(ctrl)

	c := NewController(mockPostRepo)

	assert.NotNil(t, c)
	assert.Equal(t, mockPostRepo, c.PostRepository)
}
