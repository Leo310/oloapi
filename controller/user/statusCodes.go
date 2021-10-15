package user

type statusId int16

type ustatus struct {
	StatusCode statusId
}

// itoa means that other items in const () get incremented automatically
const (
	noErr statusId = iota
	errEmailInvalid
	errEmailAlreadyRegistered
	errNameInvalid
	errPasswordInvalid
	errReviewInput
	errSomeError
	errCredentialsInvalid
	errUserNotFound
)

func (status statusId) String() string {
	return [...]string{
		"No error",
		"Email is Invalid",
		"Email is already Registered",
		"Name is Invalid",
		"Length of password should be atleast 8 and it must be a combination of uppercase letters, lowercase letters and numbers",
		"Review your Input",
		"Something went wrong, please try again later. 😕",
		"Credentials are Invalid",
		"Can not find User",
	}[status]
}

// TODO: puts enum in namespace
