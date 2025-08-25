package errors

type Code string

type Error interface {
	error
	Code() Code
	Message() string
	Previous() error
	// TODO introduce []uintptr as a type so I can add a PrettyPrint function to it
	StackTrace() []uintptr
}

type ErrorFactory interface {
	New(code Code, message string, prev error) Error
}

const (
	ErrValidation   Code = "VALIDATION"
	ErrUnauthorized Code = "UNAUTHORIZED"
	ErrConflict     Code = "CONFLICT"
	ErrDB           Code = "DB_ERROR"
	ErrExternal     Code = "EXTERNAL_ERROR"
	ErrUnknown      Code = "UNKNOWN"
)
