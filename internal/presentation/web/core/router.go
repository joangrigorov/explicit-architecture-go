package core

import (
	"app/internal/core/component/blog/application/repositories"
	"app/internal/presentation/web/core/component/blog/anonymous/v1/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(pr repositories.PostRepository) *gin.Engine {
	r := gin.Default()

	// blog routes
	{
		postController := controllers.NewPostController(pr)
		v1 := r.Group("/blogs/v1")
		v1.POST("/posts", postController.CreatePost)
		v1.GET("/posts/:id", postController.GetPost)
		v1.GET("/posts", postController.ListPosts)
		v1.DELETE("/posts/:id", postController.DeletePost)
		v1.PATCH("/posts/:id", postController.UpdatePost)
	}

	return r
}
