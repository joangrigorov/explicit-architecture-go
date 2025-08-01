package core

import (
	"app/internal/core/component/blog/application/repositories"
	"app/internal/presentation/web/core/component/blog/v1/anonymous/controllers/post"
	"app/internal/presentation/web/port/http"
)

func RegisterRoutes(
	pr repositories.PostRepository,
	r http.Router,
) {
	// blog routes
	{
		postController := post.NewController(pr)
		v1 := r.Group("/blogs/v1")
		v1.POST("/posts", postController.Create)
		v1.GET("/posts/:id", postController.GetOne)
		v1.GET("/posts", postController.Index)
		v1.DELETE("/posts/:id", postController.Delete)
		v1.PATCH("/posts/:id", postController.Update)
	}
}
