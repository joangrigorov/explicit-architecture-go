package find_user_by_id

import (
	"encoding/json"
)

type Query struct {
	ID string
}

func NewQuery(ID string) Query {
	return Query{ID: ID}
}

func (q Query) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id": q.ID,
	})
}
