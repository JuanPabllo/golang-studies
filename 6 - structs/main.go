package main

import "fmt"

type user struct {
	name     string
	age      uint8
	endereco address
}

type address struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Hello World")

	var u user

	u.name = "Juan"
	u.age = 20

	addressExample := address{"Rua 1", 10}

	user2 := user{"Ana", 19, addressExample}

	user3 := user{age: 1, name: "Luna"}
	user4 := user{name: "Luna"}

	fmt.Println(user3)
	fmt.Println(user4)
	fmt.Println(user2)

	fmt.Println(u)
}
