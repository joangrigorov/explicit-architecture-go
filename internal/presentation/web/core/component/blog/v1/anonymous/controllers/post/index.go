package post

import (
	. "app/internal/presentation/web/core/component/blog/v1/anonymous/responses"
	. "app/internal/presentation/web/core/shared_kernel/responses"
	. "app/internal/presentation/web/port/http"
	"net/http"
)

func (pc *Controller) Index(c Context) {
	posts, err := pc.PostRepository.GetAll(c.Context())

	if err != nil {
		InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, MultiPostResponse(posts))
}
