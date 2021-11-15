package user

import (
	"regexp"

	valid "github.com/asaskevich/govalidator"
)

func validUUID(uuid string) bool {
	return valid.IsUUID(uuid)
}

func validEmail(email string) bool {
	return valid.IsEmail(email)
}

func validName(name string) bool {
	return valid.IsNotNull(name) || valid.HasWhitespaceOnly(name)
}

func validPassword(password string) bool {
	re := regexp.MustCompile(`\d`) // regex check for at least one integer in string
	return len(password) >= 8 && valid.HasLowerCase(password) && valid.HasUpperCase(password) && re.MatchString(password)
}
