package errors

import "fmt"

type IdPUserNotConnected struct{}

func (i IdPUserNotConnected) Error() string {
	return "user not connected to identity provider"
}

func NewIdPUserNotConnectedError() error {
	return &IdPUserNotConnected{}
}

type IdPRequestError struct {
	err error
}

func (i IdPRequestError) Error() string {
	return fmt.Sprintf("idp request error: %s", i.err.Error())
}

func NewIdPRequestError(err error) error {
	return &IdPRequestError{err: err}
}
