package util

import (
	"regexp"

	valid "github.com/asaskevich/govalidator"
)

func ValidEmail(email string) bool {
	return valid.IsEmail(email)
}

func ValidName(name string) bool {
	return valid.IsNotNull(name) || valid.HasWhitespaceOnly(name)
}

func ValidPassword(password string) bool {
	re := regexp.MustCompile(`\d`) // regex check for at least one integer in string
	return len(password) >= 8 && valid.HasLowerCase(password) && valid.HasUpperCase(password) && re.MatchString(password)
}
