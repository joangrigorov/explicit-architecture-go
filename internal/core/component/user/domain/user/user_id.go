package user

type ID string

func (i ID) String() string {
	return string(i)
}
