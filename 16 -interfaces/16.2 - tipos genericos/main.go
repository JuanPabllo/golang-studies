package main

import "fmt"

func generica(interef interface{}) {
	fmt.Println(interef)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(1.2): true,
		"String":     1,
	}

	fmt.Println(mapa)

}
