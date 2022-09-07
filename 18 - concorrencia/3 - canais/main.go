package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// for {
	// 	msg, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(msg)
	// }

	for message := range canal {
		fmt.Println(message)
	}
	fmt.Println("Fim do programa")

}

func escrever(txt string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- txt
		time.Sleep(time.Second)
	}

	close(canal)
}
