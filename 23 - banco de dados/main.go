package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		panic(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		panic(erro)
	}

	fmt.Println("Conexão realizada com sucesso!")

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		panic(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}
