package confirm_user

import (
	"encoding/json"
)

type Command struct {
	userID string
}

func NewCommand(userID string) Command {
	return Command{userID: userID}
}

func (c Command) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userID": c.userID,
	})
}
