package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Goroutine! 3")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Goroutine! 4")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0
}

func escrever(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
