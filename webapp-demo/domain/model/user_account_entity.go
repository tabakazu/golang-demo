package model

type UserAccountEntity struct {
	ID             uint
	FamilyName     string
	GivenName      string
	Email          string
	PasswordDigest string `json:"-"`
}
