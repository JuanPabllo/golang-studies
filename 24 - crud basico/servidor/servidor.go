package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequsisao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisicao"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequsisao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o corpo da requisicao para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao preparar o statement"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	resposta := map[string]string{
		"message": "Usuário inserido com sucesso.",
		"id":      fmt.Sprintf("%d", idInserido),
		"nome":    usuario.Nome,
		"email":   usuario.Email,
	}

	respostaJSON, _ := json.Marshal(resposta)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(respostaJSON))
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select id, nome, email from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuarios"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao ler o usuario"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para json"))
		return
	}

}

func ListarUsuarioPorid(w http.ResponseWriter, r *http.Request) {

}
