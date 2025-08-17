package activities

import (
	. "app/internal/presentation/api/core/component/activity/v1/responses"
	. "app/internal/presentation/api/core/shared/responses"
	. "app/internal/presentation/api/port/http"
	"net/http"
)

func (pc *Controller) Index(c Context) {
	entries, err := pc.activityRepository.GetAll(c.Context())

	if err != nil {
		InternalServerError(c, &DefaultError{Error: "Application error."})
		return
	}
	c.JSON(http.StatusOK, Many(entries))
}
