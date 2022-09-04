package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func writer(texto string, numero ...int) {
	for _, numero := range numero {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 12, 124, 123)

	fmt.Println(totalSoma)
	writer("O número é:", 1, 2, 3, 4, 5, 6, 12, 124, 123)
}
