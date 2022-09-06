package main

import "fmt"

func closure() func() {
	text := "Dentro da função closure"

	funcao := func() {
		fmt.Println(text)
	}

	return funcao
}

func main() {
	text := "Dentro da função main"
	fmt.Println(text)

	funcaoNova := closure()
	funcaoNova()
}
