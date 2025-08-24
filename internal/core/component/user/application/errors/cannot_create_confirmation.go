package errors

type CannotCreateConfirmationError struct {
	previous error
}

func (e *CannotCreateConfirmationError) Error() string {
	return "cannot create confirmation: " + e.previous.Error()
}

func NewCannotCreateConfirmationError(previous error) *CannotCreateConfirmationError {
	return &CannotCreateConfirmationError{previous: previous}
}

