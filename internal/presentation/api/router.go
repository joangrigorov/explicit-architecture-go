package api

import (
	"app/internal/infrastructure/framework/http"
	"app/internal/presentation/api/component/activity/v1/controllers/activities"
	"app/internal/presentation/api/component/user/v1/controllers"
	"app/internal/presentation/api/shared/middleware"
)

func RegisterRoutes(
	r http.Router,
	activityController *activities.Controller,
	registrationController *controllers.RegistrationController,
) {
	api := r.Group("/api")

	// Global middleware
	api.Use(
		// We make sure only fundamentally valid JSON passes through.
		// The validation only happens for json requests.
		middleware.ValidateJSONBody,
	)

	{
		v1 := api.Group("/user/v1")

		v1.POST("/registration", registrationController.Register)
		v1.POST("/confirm", registrationController.Confirm)
	}

	// activity component public routes
	{
		v1 := api.Group("/activity/v1")

		v1.GET("/activities/:id", activityController.GetOne)
		v1.GET("/activities", activityController.Index)
	}
}
