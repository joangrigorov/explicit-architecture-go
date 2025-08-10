package activities

import (
	"app/mock/core/component/activity/application/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	activityRepository := repositories.NewMockActivityRepository(ctrl)

	c := NewController(activityRepository)

	assert.NotNil(t, c)
	assert.Equal(t, activityRepository, c.activityRepository)
}
