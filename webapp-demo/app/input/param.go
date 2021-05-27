package input

type RegisterUserAccountParam struct {
	FamilyName           string `json:"family_name"`
	GivenName            string `json:"given_name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type LoginUserAccountParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ShowUserAccountParam struct {
	ID uint `json:"id"`
}
