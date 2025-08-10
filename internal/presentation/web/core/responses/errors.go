package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Code    string `json:"code"`    // e.g. HTTP status code or app-specific code
	Message string `json:"message"` // error message text
}

func newErrorResponse(err error, code string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
}

func ApplicationError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, newErrorResponse(err, "application_error"))
}

func RequestError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, newErrorResponse(err, "invalid_request"))
}

func NotFoundError(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, newErrorResponse(err, "not_found"))
}
