package post

import (
	"app/internal/presentation/web/core/component/blog/v1/anonymous/requests"
	. "app/internal/presentation/web/core/component/blog/v1/anonymous/responses"
	. "app/internal/presentation/web/core/shared_kernel/responses"
	. "app/internal/presentation/web/port/http"
	"net/http"
)

func (pc *Controller) Update(c Context) {
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

	var req requests.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		UnprocessableEntity(c, Render(pc.translator, err, req))
		return
	}

	req.Populate(p)

	err = pc.PostRepository.Update(c.Context(), p)
	if err != nil {
		InternalServerError(c, err)
	}

	c.JSON(http.StatusOK, OnePostResponse(p))
}
