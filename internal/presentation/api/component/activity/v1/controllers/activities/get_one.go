package activities

import (
	"app/internal/core/component/activity/domain"
	. "app/internal/infrastructure/http"
	. "app/internal/presentation/api/component/activity/v1/responses"
	. "app/internal/presentation/api/shared/responses"
	"net/http"
)

func (pc *Controller) GetOne(c Context) {
	id := c.ParamString("id")

	p, err := pc.activityRepository.GetById(c.Context(), domain.ActivityID(id))
	if err != nil {
		NotFound(c, NewDefaultError(err))
		return
	}

	c.JSON(http.StatusOK, One(p))
}
