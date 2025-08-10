package activities

import (
	. "app/internal/core/component/activity/application/repositories"
)

type Controller struct {
	activityRepository ActivityRepository
}

func NewController(activityRepository ActivityRepository) *Controller {
	return &Controller{activityRepository: activityRepository}
}
