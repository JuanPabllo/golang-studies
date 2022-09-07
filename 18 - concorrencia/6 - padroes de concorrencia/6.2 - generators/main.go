package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(txt string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", txt)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal

}
