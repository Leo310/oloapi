package user

import (
	"regexp"
	"testing"
)

// Testing generate random salt
func (userenv *Userenv) TestGenerateRandomSalt(t *testing.T) {
	r, _ := regexp.Compile(`^[a-zA-Z0-9$%\?]{8}$`)
	for i := 0; i < 1000; i++ {
		salt := generateRandomSalt()
		if match := r.MatchString(salt); !match {
			t.Errorf("Error Salt: %s", salt)
		}
	}

}

func BenchmarkGenerateRandomSalt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateRandomSalt()
	}
}

// Testing user validation
// func TestUserValidation(t *testing.T) {
// 	inputoutput := []struct {

// 	}
// }
