package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"nome":      "Juan",
		"sobrenome": "Pablo",
	}

	fmt.Println(user["nome"])

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pereira",
		},
		"curso": {
			"nome": "Engenharia",
		},
	}

	fmt.Println(user2)

	delete(user2, "nome")
	fmt.Println(user2)

	user2["idade"] = map[string]string{
		"numero": "30",
	}
	fmt.Println(user2)

}
