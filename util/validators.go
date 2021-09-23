package util

import (
	"oloapi/models"
	"regexp"

	valid "github.com/asaskevich/govalidator"
)

// ValidateRegister func validates the body of user for registration
func ValidateRegister(u *models.User) *models.UserErrors {
	e := &models.UserErrors{}

	if !valid.IsEmail(u.Email) {
		e.Err, e.Email = true, "Must be a valid email"
	}

	re := regexp.MustCompile("\\d") // regex check for at least one integer in string
	if !(len(u.Password) >= 8 && valid.HasLowerCase(u.Password) && valid.HasUpperCase(u.Password) && re.MatchString(u.Password)) {
		e.Err, e.Password = true, "Length of password should be atleast 8 and it must be a combination of uppercase letters, lowercase letters and numbers"
	}

	return e
}
