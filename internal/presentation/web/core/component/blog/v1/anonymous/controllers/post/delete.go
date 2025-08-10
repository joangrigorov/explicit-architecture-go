package post

import (
	. "app/internal/presentation/web/core/responses"
	. "app/internal/presentation/web/port/http"
	"strconv"
)

func (pc *Controller) Delete(c Context) {
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

	err = pc.PostRepository.Delete(c.Context(), p)
	if err != nil {
		InternalServerError(c, err)
	}

	c.NoContent()
}
