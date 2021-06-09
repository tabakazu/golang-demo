package main

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type User struct {
	Email string `validate:"required,email,unique"`
}

func main() {
	u := &User{Email: "example__example.com"}
	validateUser(u)

	u2 := &User{Email: ""}
	validateUser(u2)

	u3 := &User{Email: "unique@example.com"}
	validateUser(u3)
}

func validateUser(u *User) {
	validate := validator.New()
	validate.RegisterValidation("unique", ValidateUserEmailUnique)

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

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
			return &User{"unique@example.com"}
		}
		return nil
	}

	u := findByEmail(fl.Field().String())
	return u == nil
}
