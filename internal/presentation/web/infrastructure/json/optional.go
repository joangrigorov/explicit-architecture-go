package json

import (
	"encoding/json"
	"fmt"
)

type Optional[T any] struct {
	IsSet bool
	Value *T
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	o.IsSet = true
	if string(data) == "null" {
		o.Value = nil
		return nil
	}
	var val T
	if err := json.Unmarshal(data, &val); err != nil {
		return fmt.Errorf("optional: %w", err)
	}
	o.Value = &val
	return nil
}
