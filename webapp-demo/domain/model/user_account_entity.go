package model

type UserAccountEntity struct {
	ID             uint
	FamilyName     string `db:"family_name"`
	GivenName      string `db:"given_name"`
	Email          string
	PasswordDigest string `db:"password_digest" json:"-"`
}
