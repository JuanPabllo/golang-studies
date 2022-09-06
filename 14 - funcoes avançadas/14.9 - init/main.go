package main

import "fmt"

var n int

func init() {
	fmt.Println("Funcao init sendo executada!")
	n = 13
}

func main() {
	fmt.Println("Funcao main sendo executada!")
	fmt.Println(n)
}
