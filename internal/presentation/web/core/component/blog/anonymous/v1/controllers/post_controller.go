package controllers

import (
	"app/internal/core/component/blog/application/repositories"
	"app/internal/presentation/web/core/component/blog/anonymous/v1/requests"
	"app/internal/presentation/web/core/component/blog/anonymous/v1/responses"
	re "app/internal/presentation/web/core/shared_kernel"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostController struct {
	PostRepository repositories.PostRepository
}

func (pc *PostController) ListPosts(c *gin.Context) {
	posts, err := pc.PostRepository.GetAll(c.Request.Context())

	if err != nil {
		re.ApplicationError(c, err)
		return
	}

	c.JSON(http.StatusOK, responses.MultiPostResponse(posts))
}

func (pc *PostController) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		re.RequestError(c, err)
		return
	}

	p, err := pc.PostRepository.GetById(c.Request.Context(), id)
	if err != nil {
		re.NotFoundError(c, err)
		return
	}

	c.JSON(http.StatusOK, responses.OnePostResponse(p))
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var req requests.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		re.RequestError(c, err)
		return
	}

	err := pc.PostRepository.Create(c.Request.Context(), req.ToPost())
	if err != nil {
		re.ApplicationError(c, err)
	}
}

func (pc *PostController) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		re.RequestError(c, err)
		return
	}

	p, err := pc.PostRepository.GetById(c.Request.Context(), id)
	if err != nil {
		re.NotFoundError(c, err)
		return
	}

	err = pc.PostRepository.Delete(c.Request.Context(), p)
	if err != nil {
		re.ApplicationError(c, err)
	}

	c.Status(http.StatusNoContent)
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		re.RequestError(c, err)
		return
	}

	p, err := pc.PostRepository.GetById(c.Request.Context(), id)
	if err != nil {
		re.NotFoundError(c, err)
		return
	}

	var req requests.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		re.RequestError(c, err)
		return
	}

	req.Populate(p)

	err = pc.PostRepository.Update(c.Request.Context(), p)
	if err != nil {
		re.ApplicationError(c, err)
	}

	c.JSON(http.StatusOK, responses.OnePostResponse(p))
}

func NewPostController(postRepository repositories.PostRepository) *PostController {
	return &PostController{PostRepository: postRepository}
}
