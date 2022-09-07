package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!") // goroutine
	escrever("Programando em Go!")
}

func escrever(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
