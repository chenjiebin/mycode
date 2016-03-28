// https://github.com/go-validator/validator

package main

import (
	"fmt"

	"gopkg.in/validator.v2"
)

func main() {
	type NewUserRequest struct {
		Username string `validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
		Name     string `validate:"nonzero"`
		Age      int    `validate:"min=21"`
		Password string `validate:"min=8"`
	}

	nur := NewUserRequest{Username: "something", Name: "Name", Age: 22, Password: "12345678"}
	if err := validator.Validate(nur); err != nil {
		fmt.Println(err)
	}
	fmt.Println(nur)
}
