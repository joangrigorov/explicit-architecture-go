package requests

import (
	"app/internal/core/component/user/application/commands/register_user"
)

type Registration struct {
	Email     string `json:"email" binding:"required,email"`
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

func (r *Registration) NewRegisterUserCommand(userID string) register_user.Command {
	return register_user.NewCommand(
		userID,
		r.Username,
		r.Email,
		r.FirstName,
		r.LastName,
	)
}
