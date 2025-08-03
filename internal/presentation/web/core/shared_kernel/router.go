package shared_kernel

import (
	"app/internal/presentation/web/core/component/blog/v1/anonymous/controllers/post"
	"app/internal/presentation/web/core/shared_kernel/middleware"
	"app/internal/presentation/web/port/http"
)

func RegisterRoutes(
	r http.Router,
	postController *post.Controller,
) {
	// Global middleware
	r.Use(
		// We make sure only fundamentally valid JSON passes through.
		// The validation only happens for json requests.
		middleware.ValidateJSONBody,
		middleware.ValidateJSONBody,
	)

	// blog component routes
	{
		v1 := r.Group("/blogs/v1")

		v1.POST("/posts", middleware.RequiresJSON, postController.Create)
		v1.PATCH("/posts/:id", middleware.RequiresJSON, postController.Update)
		v1.DELETE("/posts/:id", postController.Delete)

		v1.GET("/posts/:id", postController.GetOne)
		v1.GET("/posts", postController.Index)
	}
}
