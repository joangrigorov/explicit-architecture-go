package uuid

import "github.com/google/uuid"

func Parse[T ~string](id T) uuid.UUID {
	parsedId, err := uuid.Parse(string(id))

	if err != nil {
		panic("invalid uuid")
	}

	return parsedId
}
