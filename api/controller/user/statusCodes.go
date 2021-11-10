package user

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

	errEmailAlreadyRegistered = "EMAIL_ALREADY_REGISTERED"
	errCredentialsInvalid     = "CREDENTIALS_INVALID"
	errUserNotFound           = "USER_NOT_FOUND"
	errLocationNotFound       = "LOCATION_NOT_FOUND"

	errReviewInput     = "REVIEW_INPUT"
	errEmailInvalid    = "EMAIL_INVALID"
	errNameInvalid     = "NAME_INVALID"
	errPasswordInvalid = "PASSWORD_INVALID"
)

// TODO: puts enum in namespace
