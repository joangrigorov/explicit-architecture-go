package post

import (
	"app/internal/presentation/web/core/component/blog/v1/anonymous/requests"
	. "app/internal/presentation/web/core/component/blog/v1/anonymous/responses"
	. "app/internal/presentation/web/core/shared_kernel/responses"
	. "app/internal/presentation/web/port/http"
	"net/http"
)

func (pc *Controller) Create(c Context) {
	var req requests.CreatePost
	if err := c.ShouldBindJSON(&req); err != nil {
		UnprocessableEntity(c, Render(pc.translator, err, req))
		return
	}

	post := req.ToPost()
	err := pc.PostRepository.Create(c.Context(), post)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, OnePostResponse(post))
}
