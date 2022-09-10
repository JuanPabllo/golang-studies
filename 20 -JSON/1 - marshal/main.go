package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Nome  string `json:"nome"` // tag
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorro := dog{"Rex", "Dálmata", 3}
	fmt.Println(cachorro)

	cachorroEmJSON, erro := json.Marshal(cachorro)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	jsonFormatted := bytes.NewBuffer(cachorroEmJSON)
	fmt.Println(jsonFormatted)

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Vira-lata",
	}

	cachorro2emJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2emJSON)
	fmt.Println(bytes.NewBuffer(cachorro2emJSON))
}
