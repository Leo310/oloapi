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
	errSomeError

	errEmailAlreadyRegistered
	errCredentialsInvalid
	errUserNotFound
	errLocationNotFound

	errReviewInput
	errEmailInvalid
	errNameInvalid
	errPasswordInvalid
)

func (status statusId) String() string {
	return [...]string{
		"No error",
		"Internal Server error",
		"Something went wrong, please try again later. 😕",

		"Email is already Registered",
		"Credentials are Invalid",
		"Can not find User",
		"Could not find Location",

		"Review your Input",
		"Email is Invalid",
		"Name is Invalid",
		"Length of password should be atleast 8 and it must be a combination of uppercase letters, lowercase letters and numbers",
	}[status]
}

// TODO: puts enum in namespace
