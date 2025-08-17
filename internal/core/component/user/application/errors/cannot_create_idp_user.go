package errors

import "fmt"

type CannotCreateIdPUserError struct {
	err error
}

func (u *CannotCreateIdPUserError) Error() string {
	return fmt.Sprintf("cannot create user at identity provider: %v", u.err)
}

func NewCannotCreateIdPUserError(err error) error {
	return &CannotCreateIdPUserError{err: err}
}
