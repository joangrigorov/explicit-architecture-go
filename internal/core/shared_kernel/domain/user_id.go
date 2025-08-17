package domain

type UserID string

func (i UserID) String() string {
	return string(i)
}
