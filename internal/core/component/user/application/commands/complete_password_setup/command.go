package complete_password_setup

import (
	"encoding/json"
)

type Command struct {
	userID   string
	password string
}

func (c Command) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userId":   c.userID,
		"password": "_",
	})
}

func NewCommand(userID string, password string) Command {
	return Command{userID: userID, password: password}
}
