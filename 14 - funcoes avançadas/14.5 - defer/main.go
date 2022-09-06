package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func alunoAprovado(nota1, nota2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será calculado.")
	fmt.Println("Entrando na funcao para verificar se aluno esta aprovado")

	media := (nota1 + nota2) / 2

	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() {

	fmt.Println(alunoAprovado(5, 6))
}
