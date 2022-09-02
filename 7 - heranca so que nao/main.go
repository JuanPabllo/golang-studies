package main

import "fmt"

type person struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type student struct {
	person
	curso     string
	faculdade string
}

func main() {
	pessoa1 := person{"João", "Silva", 20, 180}

	estudante1 := student{pessoa1, "Engenharia", "USP"}

	fmt.Println(pessoa1)
	fmt.Println(estudante1.nome)
}
