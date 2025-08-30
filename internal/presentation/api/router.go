package api

import (
	"app/internal/infrastructure/framework/http"
	"app/internal/presentation/api/component/activity/v1/controllers/activities"
	uc "app/internal/presentation/api/component/user/v1/controllers"
	"app/internal/presentation/api/shared/middleware"
)

func RegisterRoutes(
	r http.Router,
	activityController *activities.Controller,
	registrationController *uc.Registration,
	verificationController *uc.Verification,
) {
	// Global middleware
	r.Use(
		// We make sure only fundamentally valid JSON passes through.
		// The validation only happens for json requests.
		middleware.ValidateJSONBody,
		middleware.ResponseContentTypeJSON,
	)

	{
		v1 := r.Group("/user/v1")

		v1.POST("/registration", registrationController.Register)
		v1.GET("/verifications/:id/preflight", verificationController.PreflightValidate)
		v1.POST("/verifications/:id/password-setup", verificationController.PasswordSetup)
	}

	// activity component public routes
	{
		v1 := r.Group("/activity/v1")

		v1.GET("/activities/:id", activityController.GetOne)
		v1.GET("/activities", activityController.Index)
	}
}
