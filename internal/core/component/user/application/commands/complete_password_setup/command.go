package complete_password_setup

import (
	"app/internal/core/component/user/domain/user"
	"encoding/json"
)

type Command struct {
	userID   user.ID
	password string
}

func (c Command) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userId":   c.userID,
		"password": "_",
	})
}

func NewCommand(
	userID user.ID,
	password string,
) Command {
	return Command{
		userID:   userID,
		password: password,
	}
}
