package post

import (
	. "app/internal/presentation/web/core/component/blog/v1/anonymous/responses"
	. "app/internal/presentation/web/core/shared_kernel/responses"
	. "app/internal/presentation/web/port/http"
	"net/http"
)

func (pc *Controller) GetOne(c Context) {
	id, err := c.ParamInt("id")
	if err != nil {
		BadRequest(c, NewDefaultError(err))
		return
	}

	p, err := pc.PostRepository.GetById(c.Context(), id)
	if err != nil {
		NotFound(c, NewDefaultError(err))
		return
	}

	c.JSON(http.StatusOK, OnePostResponse(p))
}
