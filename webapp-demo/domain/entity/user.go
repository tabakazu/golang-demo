package entity

import (
	"time"

	"github.com/tabakazu/webapi-app/domain/value"
)

type User struct {
	GivenName      string
	FamilyName     string
	BirthedOn      time.Time
	PasswordDigest value.PasswordDigest
}

func (u *User) FullName() string {
	return u.GivenName + u.FamilyName
}

func (u *User) IsCorrectPassword(pass value.Password) bool {
	return u.PasswordDigest.IsCorrectPassword(pass)
}
