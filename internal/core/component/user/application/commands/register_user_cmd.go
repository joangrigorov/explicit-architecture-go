package commands

import (
	. "app/internal/core/component/user/domain"
	. "app/internal/core/shared_kernel/domain"
	"encoding/json"
)

type RegisterUserCommand struct {
	userID    string
	username  string
	password  string
	email     string
	firstName string
	lastName  string
}

func (r RegisterUserCommand) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userID":    r.userID,
		"username":  r.username,
		"email":     r.email,
		"firstName": r.firstName,
		"lastName":  r.lastName,
	})
}

func NewRegisterUserCommand(
	userID string,
	username string,
	password string,
	email string,
	firstName string,
	lastName string,
) RegisterUserCommand {
	return RegisterUserCommand{
		userID:    userID,
		username:  username,
		password:  password,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
	}
}
