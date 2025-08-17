package requests

type ConfirmationRequest struct {
	UserID string `json:"user_id" binding:"required"`
}
