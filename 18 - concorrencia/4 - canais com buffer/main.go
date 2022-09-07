package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	msg := <-canal
	msg2 := <-canal

	fmt.Println(msg)
	fmt.Println(msg2)
}
