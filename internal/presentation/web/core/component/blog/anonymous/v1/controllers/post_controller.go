package controllers

import (
	"app/internal/core/component/blog/application/repositories"
	"app/internal/presentation/web/core/component/blog/anonymous/v1/requests"
	"app/internal/presentation/web/responses"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostRepository repositories.PostRepository
}

func NewPostController(postRepository repositories.PostRepository) *PostController {
	return &PostController{PostRepository: postRepository}
}

func (pc *PostController) ListPosts(c *gin.Context) {

}

func (pc *PostController) GetPost(c *gin.Context) {

}

func (pc *PostController) CreatePost(c *gin.Context) {
	var req requests.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, responses.NewErrorResponse(err, "request_error"))
		return
	}

	post := req.ToPost()
	err := pc.PostRepository.Create(c.Request.Context(), &post)
	if err != nil {
		c.JSON(500, responses.NewErrorResponse(err, "application_error"))
	}
}
