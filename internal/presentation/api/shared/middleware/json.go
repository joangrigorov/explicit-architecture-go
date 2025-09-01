package middleware

import (
	ctx "app/internal/infrastructure/framework/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateJSONBody middleware checks if the format of the JSON body is valid.
// Check is performed on request bodies only when the request content type
// is application/json
func ValidateJSONBody(c ctx.Context) {
	if c.IsPost() || c.IsPut() || c.IsPatch() {
		if !c.IsJsonRequest() {
			c.Next()
			return
		}

		if !c.IsJsonBodyValid() {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
			return
		}
	}

	c.Next()
}

func ResponseContentTypeJSON(c ctx.Context) {
	c.SetResponseHeader("Content-Type", "application/json; charset=utf-8")
	c.Next()
}

// RequiresJSON middleware ensures that only requests with content type
// application/json continue.
func RequiresJSON(c ctx.Context) {
	if !c.IsJsonRequest() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "This API only serves JSON requests"})
		return
	}

	c.Next()
}
