package activities

import (
	"app/internal/core/component/activity/domain"
	. "app/internal/presentation/web/core/component/activity/v1/responses"
	. "app/internal/presentation/web/core/shared/responses"
	. "app/internal/presentation/web/port/http"
	"net/http"
)

func (pc *Controller) GetOne(c Context) {
	id := c.ParamString("id")

	p, err := pc.activityRepository.GetById(c.Context(), domain.ActivityId(id))
	if err != nil {
		NotFound(c, NewDefaultError(err))
		return
	}

	c.JSON(http.StatusOK, One(p))
}
