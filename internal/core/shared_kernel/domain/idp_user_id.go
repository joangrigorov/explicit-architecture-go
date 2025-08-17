package domain

type IdPUserId string

func (i IdPUserId) String() string {
	return string(i)
}
