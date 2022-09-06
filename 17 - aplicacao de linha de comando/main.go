package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicação := app.Gerar()
	if erro := aplicação.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
