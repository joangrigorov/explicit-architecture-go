package shared_kernel

import (
	"app/internal/presentation/web/core/component/activity/v1/controllers/activities"
	"app/internal/presentation/web/core/shared_kernel/middleware"
	"app/internal/presentation/web/port/http"
)

func RegisterRoutes(
	r http.Router,
	activityController *activities.Controller,
) {
	// Global middleware
	r.Use(
		// We make sure only fundamentally valid JSON passes through.
		// The validation only happens for json requests.
		middleware.ValidateJSONBody,
	)

	// activity component public routes
	{
		v1 := r.Group("/activity/v1")
		//
		//v1.POST("/activities", middleware.RequiresJSON, activityController.Create)
		//v1.PATCH("/posts/:id", middleware.RequiresJSON, activityController.Update)
		//v1.DELETE("/posts/:id", activityController.Delete)

		v1.GET("/activities/:id", activityController.GetOne)
		v1.GET("/activities", activityController.Index)
	}
}
