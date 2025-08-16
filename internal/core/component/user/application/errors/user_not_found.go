package errors

type UserNotFoundError struct{}

func (u *UserNotFoundError) Error() string {
	return "user not found"
}

func NewUserNotFoundError() error {
	return &UserNotFoundError{}
}
