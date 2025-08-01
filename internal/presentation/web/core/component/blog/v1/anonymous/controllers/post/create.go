package post

import (
	"app/internal/presentation/web/core/component/blog/v1/anonymous/requests"
	. "app/internal/presentation/web/core/responses"
	. "app/internal/presentation/web/port/http"
)

func (pc *Controller) Create(c Context) {
	var req requests.CreatePost
	if err := c.ShouldBindJSON(&req); err != nil {
		UnprocessableEntity(c, err, req)
		return
	}

	post := req.ToPost()
	err := pc.PostRepository.Create(c.Context(), post)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	c.JSON(201, post)
}
