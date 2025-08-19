package responses

import (
	"app/internal/core/component/user/domain"
	"time"
)

type RegistrationResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Confirmed bool      `json:"confirmed"`
}

func NewRegistrationResponse(user *domain.User) *RegistrationResponse {
	return &RegistrationResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Confirmed: false,
	}
}
