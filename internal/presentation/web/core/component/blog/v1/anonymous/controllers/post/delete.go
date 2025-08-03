package post

import (
	. "app/internal/presentation/web/core/shared_kernel/responses"
	. "app/internal/presentation/web/port/http"
)

func (pc *Controller) Delete(c Context) {
	id, err := c.ParamInt("id")
	if err != nil {
		BadRequest(c, NewDefaultError(err))
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
