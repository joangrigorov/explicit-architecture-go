package activities

import (
	. "app/internal/infrastructure/http"
	. "app/internal/presentation/api/component/activity/v1/responses"
	. "app/internal/presentation/api/shared/responses"
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
