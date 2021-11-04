package user

type statusId int16

// TODO statusCode lower case
type ustatus struct {
	StatusCode statusId
}

// itoa means that other items in const () get incremented automatically
const (
	noErr statusId = iota
	errServerInternal
	errEmailInvalid
	errEmailAlreadyRegistered
	errNameInvalid
	errPasswordInvalid
	errReviewInput
	errSomeError
	errCredentialsInvalid
	errUserNotFound
	errLocationNotFound
)

func (status statusId) String() string {
	return [...]string{
		"No error",
		"Internal Server error",
		"Email is Invalid",
		"Email is already Registered",
		"Name is Invalid",
		"Length of password should be atleast 8 and it must be a combination of uppercase letters, lowercase letters and numbers",
		"Review your Input",
		"Something went wrong, please try again later. 😕",
		"Credentials are Invalid",
		"Can not find User",
		"Could not find Location",
	}[status]
}

// TODO: puts enum in namespace
