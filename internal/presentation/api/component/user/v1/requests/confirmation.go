package requests

type Confirmation struct {
	UserID string `json:"user_id" binding:"required"`
}
