package main

import "fmt"

type user struct {
	name  string
	idade uint8
}

func (u user) save() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.name)
}

func (u user) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *user) fazerAniversario() {
	u.idade++
}

func main() {
	user1 := user{"Juan", 20}
	user1.save()

	user2 := user{name: "Ana", idade: 18}
	user2.save()

	user3 := user{"Luna", 1}

	fmt.Println(user3.maiorDeIdade())

	user2.fazerAniversario()
	fmt.Println(user2.idade)
}
