package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "rodovia", "estrada"}
	enderecoEmLetrasMinusculas := strings.ToLower(endereco)

	primeiraPalavra := strings.Split(enderecoEmLetrasMinusculas, " ")[0]

	enderecoTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTipoValido = true
		}
	}

	if enderecoTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "tipo inválido"
}
