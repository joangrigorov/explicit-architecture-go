package user

type Username string

func (u Username) IsValid() bool {
	return u.String() != ""
}

func (u Username) String() string {
	return string(u)
}

func NewUsername(username string) (Username, error) {
	u := Username(username)
	if !u.IsValid() {
		return "", InvalidUsernameError{}
	}
	return u, nil
}

type InvalidUsernameError struct{}

func (i InvalidUsernameError) Error() string {
	return "invalid username provided"
}
