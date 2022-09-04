package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controle")

	number := -5

	if number > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Número é maior que zero")
	} else if number > -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Entre -10 e 0")
	}

}
