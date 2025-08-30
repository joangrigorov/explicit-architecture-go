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
	ErrValidation         Code = "VALIDATION"
	ErrUnauthorized       Code = "UNAUTHORIZED"
	ErrConflict           Code = "CONFLICT"
	ErrDB                 Code = "DB_ERROR"
	ErrNotFound           Code = "NOT_FOUND"
	ErrExternal           Code = "EXTERNAL_ERROR"
	ErrTypeMismatch       Code = "TYPE_MISMATCH"
	ErrCommandHandling    Code = "COMMAND_HANDLING_ERROR"
	ErrQueryHandlingError Code = "QUERY_HANDLING_ERROR"
	ErrMail               Code = "MAIL_ERROR"
	ErrUnknown            Code = "UNKNOWN"
)
