package post

import (
	. "app/internal/presentation/web/core/component/blog/v1/anonymous/responses"
	. "app/internal/presentation/web/core/responses"
	. "app/internal/presentation/web/port/http"
	"net/http"
	"strconv"
)

func (pc *Controller) GetOne(c Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		BadRequest(c, err)
		return
	}

	p, err := pc.PostRepository.GetById(c.Context(), id)
	if err != nil {
		NotFound(c, err)
		return
	}

	c.JSON(http.StatusOK, OnePostResponse(p))
}
