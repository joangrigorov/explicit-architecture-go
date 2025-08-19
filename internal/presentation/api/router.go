package api

import (
	"app/internal/infrastructure/http"
	"app/internal/presentation/api/component/activity/v1/controllers/activities"
	"app/internal/presentation/api/component/user/v1/controllers"
	"app/internal/presentation/api/shared/middleware"
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
