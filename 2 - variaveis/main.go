package main

import "fmt"

func main() {
	var variavel1 string = "Teste"
	variavel := "Teste 2"

	var (
		variavel2 string = "Teste 3"
		variavel3 string = "Teste 4"
	)

	variave4, variavel5 := "Teste 5", "Teste 6"

	const constante1 string = "Teste 7"

	fmt.Println(variavel1)
	fmt.Println(variavel)
	fmt.Println(variavel2, variavel3)
	fmt.Println(variave4, variavel5)
	fmt.Println(constante1)
}
