package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("avenida dos imigrantes")

	fmt.Println(tipoEndereco)
}
