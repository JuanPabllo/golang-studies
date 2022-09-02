package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 10000
	var numero2 uint16 = 10000

	// alias
	// int32 = rune
	// byte = uint8

	var numero3 rune = 1234
	var numero4 byte = 123

	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 121231232133.45
	numeroReal3 := 123.45

	var str string = "Texto"
	str2 := "Texto1"

	char := 'B'

	var textoVazio int16

	var booleano1 bool = true

	var erro error = errors.New("Erro interno")

	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numeroReal1)
	fmt.Println(numeroReal2)
	fmt.Println(numeroReal3)
	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(char)
	fmt.Println(textoVazio)
	fmt.Println(booleano1)
	fmt.Println(erro)
}
