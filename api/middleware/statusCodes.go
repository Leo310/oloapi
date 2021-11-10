package middleware

type statusCode string

// TODO statusCode lower case
type ustatus struct {
	StatusCode statusCode
}

// itoa means that other items in const () get incremented automatically
const (
	noErr             statusCode = "SOBER"
	errServerInternal            = "INTERNAL"
	errSomeError                 = "SOME_ERROR"

	errTokenUnavailable      = "TOKEN_UNAVAILABLE"
	errTokenExpired          = "TOKEN_EXPIRED"
	errTokenNotActiveExpired = "TOKEN_NOT_ACTIVE_EXPIRED"
)

// TODO: puts enum in namespace
