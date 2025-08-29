package find_user_by_id

import (
	"encoding/json"
)

type Query struct {
	ID string
}

func (q Query) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id": q.ID,
	})
}
