package errors

type Code string

type StackTrace []uintptr

type Error interface {
	error
	Code() Code
	Message() string
	Previous() error
	StackTrace() StackTrace
	PrettyPrint() []string
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
