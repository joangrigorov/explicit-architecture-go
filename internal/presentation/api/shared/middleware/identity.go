package middleware

import "app/internal/infrastructure/framework/http"

func RequiresIdentity(c http.Context) {
	userID := c.GetHeader("x-app-user-id")

	if userID == "" {
		c.AbortWithStatusJSON(401, map[string]interface{}{
			"error": "unauthorized",
		})
		return
	}
}
