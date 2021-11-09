package middleware

type statusId int16

// TODO statusCode lower case
type ustatus struct {
	StatusCode statusId
}

// itoa means that other items in const () get incremented automatically
const (
	noErr statusId = iota
	errServerInternal
	errTokenUnavailable
	errTokenExpired
	errTokenNotActiveExpired
)

func (status statusId) String() string {
	return [...]string{
		"No error",
		"Internal Server error",
		"Token is unavailable",
		"Token is expired",
		"Token is either not active yet or expired",
	}[status]
}

// TODO: puts enum in namespace
