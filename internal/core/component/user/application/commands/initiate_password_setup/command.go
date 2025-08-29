package initiate_password_setup

import (
	"encoding/json"
)

type Command struct {
	userID      string
	senderEmail string
}

func NewCommand(userID string, senderEmail string) Command {
	return Command{userID: userID, senderEmail: senderEmail}
}

func (s Command) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"userID":      s.userID,
		"senderEmail": s.senderEmail,
	})
}
