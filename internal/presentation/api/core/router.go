package core

import (
	"app/internal/presentation/api/core/component/activity/v1/controllers/activities"
	"app/internal/presentation/api/core/component/user/v1/controllers"
	"app/internal/presentation/api/core/shared/middleware"
	"app/internal/presentation/api/port/http"
)

func RegisterRoutes(
	r http.Router,
	activityController *activities.Controller,
	registrationController *controllers.RegistrationController,
) {
	// Global middleware
	r.Use(
		// We make sure only fundamentally valid JSON passes through.
		// The validation only happens for json requests.
		middleware.ValidateJSONBody,
	)

	{
		v1 := r.Group("/user/v1")

		v1.POST("/registration", registrationController.Register)
		v1.POST("/confirm", registrationController.Confirm)
	}

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
