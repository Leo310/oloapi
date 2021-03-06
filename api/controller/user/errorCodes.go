package user

type errorCode string

// TODO statusCode lower case
type uerror struct {
	ErrorCode errorCode
}

// itoa means that other items in const () get incremented automatically
const (
	errServerInternal errorCode = "INTERNAL"
	errSomeError                = "SOME_ERROR"

	errEmailAlreadyRegistered = "EMAIL_ALREADY_REGISTERED"
	errCredentialsInvalid     = "CREDENTIALS_INVALID"
	errUserNotFound           = "USER_NOT_FOUND"
	errLocationNotFound       = "LOCATION_NOT_FOUND"

	errReviewInput     = "REVIEW_INPUT"
	errEmailInvalid    = "EMAIL_INVALID"
	errNameInvalid     = "NAME_INVALID"
	errPasswordInvalid = "PASSWORD_INVALID"

	errTokenUnavailable       = "TOKEN_UNAVAILABLE"
	errTokenExpired           = "TOKEN_EXPIRED"
	errTokenNotActiveExpired  = "TOKEN_NOT_ACTIVE_EXPIRED"
	errTokenBad               = "TOKEN_INVALID"
	errTokenOfNonexistingUser = "TOKEN_OF_NONEXISTING_USER"
)

// TODO: puts enum in namespace
