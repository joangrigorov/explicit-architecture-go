package user

type IdPUserID string

func (i IdPUserID) String() string {
	return string(i)
}
