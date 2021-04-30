package value

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordDigest string

func (d PasswordDigest) IsCorrectPassword(pass Password) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(d), []byte(pass)); err != nil {
		return false
	}
	return true
}
