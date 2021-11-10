package middleware

type errorCode string

// TODO statusCode lower case
type merror struct {
	ErrorCode errorCode
}

// itoa means that other items in const () get incremented automatically
const (
	errServerInternal errorCode = "INTERNAL"
	errSomeError                = "SOME_ERROR"

	errTokenUnavailable      = "TOKEN_UNAVAILABLE"
	errTokenExpired          = "TOKEN_EXPIRED"
	errTokenNotActiveExpired = "TOKEN_NOT_ACTIVE_EXPIRED"
)

// TODO: puts enum in namespace
