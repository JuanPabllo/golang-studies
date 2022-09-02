package main

import "fmt"

func main() {
	// ARiTIMÉTICOS
	sum := 1 + 2
	sub := 2 - 1
	mult := 2 * 2
	div := 4 / 2
	mod := 10 % 2

	fmt.Println(sum, sub, mult, div, mod)

	// ATRIBUIÇÃO

	var text string
	text = "Texto 1"
	fmt.Println(text)

	// RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// LÓGICOS
	// AND
	fmt.Println(true && true)
	fmt.Println(true && false)

	// OR
	fmt.Println(true || true)
	fmt.Println(true || false)

	// NOT
	fmt.Println(!true)
	fmt.Println(!false)

	// UNÁRIO
	num := 10
	num++
	num += 15
	fmt.Println(num)

	// TERNÁRIO
	// Não existe no Go

	// BITWISE
	// &  |  ^  &^
	fmt.Println("01 -", 10&11)
	fmt.Println("02 -", 10|11)
	fmt.Println("03 -", 10^11)
	fmt.Println("04 -", 10&^11)

	// DESLOCAMENTO DE BITS
	fmt.Println("05 -", 10<<1)
	fmt.Println("06 -", 10>>1)

}
