package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Increment i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Maria"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
