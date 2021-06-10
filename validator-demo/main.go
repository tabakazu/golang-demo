package main

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type User struct {
	Email    string `validate:"required,email,unique"`
	Password string `validate:"required,password"`
}

func main() {
	u := &User{Email: "example__example.com", Password: "foobarbaz"}
	validateUser(u)

	u2 := &User{Email: "", Password: ""}
	validateUser(u2)

	u3 := &User{Email: "unique@example.com", Password: "Foobarbaz123!"}
	validateUser(u3)
}

func validateUser(u *User) {
	validate := validator.New()
	validate.RegisterValidation("unique", ValidateUserEmailUnique)
	validate.RegisterValidation("password", ValidatePassword)

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTranslation("password", trans, func(ut ut.Translator) error {
		return ut.Add("password", "{0} must be a valid password", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("password", fe.Field())
		return t
	})

	if err := validate.Struct(u); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			fmt.Println(e.Translate(trans))
		}
	}
}

func ValidateUserEmailUnique(fl validator.FieldLevel) bool {
	// db 検索想定
	findByEmail := func(email string) *User {
		if email == "unique@example.com" {
			return &User{"unique@example.com", ""}
		}
		return nil
	}

	u := findByEmail(fl.Field().String())
	return u == nil
}

func ValidatePassword(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 12
}
