package value

import (
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p Password) GenerateDigest() (PasswordDigest, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return PasswordDigest(""), err
	}
	return PasswordDigest(hash), nil
}

func (p Password) ForceGenerateDigest() PasswordDigest {
	digest, _ := p.GenerateDigest()
	return digest
}
