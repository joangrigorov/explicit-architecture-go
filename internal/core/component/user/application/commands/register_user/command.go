package register_user

import (
	"encoding/json"
)

type Command struct {
	userID    string
	username  string
	email     string
	firstName string
	lastName  string
}

func NewCommand(
	userID string,
	username string,
	email string,
	firstName string,
	lastName string,
) Command {
	return Command{
		userID:    userID,
		username:  username,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
	}
}

func (r Command) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userID":    r.userID,
		"username":  r.username,
		"email":     r.email,
		"firstName": r.firstName,
		"lastName":  r.lastName,
	})
}
