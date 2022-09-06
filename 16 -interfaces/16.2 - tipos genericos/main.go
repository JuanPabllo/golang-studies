package main

import "fmt"

func generica(interef interface{}) {
	fmt.Println(interef)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

}
