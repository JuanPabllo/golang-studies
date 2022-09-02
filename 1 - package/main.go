package main

import (
	"fmt"
	"modulo/helper"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Teste 1 ")
	helper.Writer()

	fmt.Println(validateEmail("1"))
}

func validateEmail(email string) error {
	err := checkmail.ValidateFormat(email)
	return err
}
