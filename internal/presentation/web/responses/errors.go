package responses

type ErrorResponse struct {
	Code    string `json:"code"`    // e.g. HTTP status code or app-specific code
	Message string `json:"message"` // error message text
}

func NewErrorResponse(err error, code string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
}
