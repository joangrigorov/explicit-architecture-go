package requests

import (
	"app/internal/core/component/user/application/commands"
)

type Registration struct {
	Email     string `json:"email" binding:"required,email"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

func (r *Registration) NewRegisterUserCommand(userID string) commands.RegisterUserCommand {
	return commands.NewRegisterUserCommand(
		userID,
		r.Username,
		r.Password,
		r.Email,
		r.FirstName,
		r.LastName,
	)
}
