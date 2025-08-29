package find_user_by_id

import "time"

type UserDTO struct {
	ID          string     `json:"id"`
	Email       string     `json:"email"`
	Username    string     `json:"username"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Role        string     `json:"role"`
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`
	IdPUserID   *string    `json:"idp_user_id,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
