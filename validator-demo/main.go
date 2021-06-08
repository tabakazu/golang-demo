package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	type S struct {
		Email string `validate:"required,email"`
	}

	validate := validator.New()

	s1 := &S{Email: "example@example.com"}
	if err := validate.Struct(s1); err != nil {
		fmt.Printf("s1 error : %v\n", err)
	}

	s2 := &S{Email: "example_example.com"}
	if err := validate.Struct(s2); err != nil {
		fmt.Printf("s2 error : %v\n", err)
	}
}
