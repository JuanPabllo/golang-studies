package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.ListarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.ListarUsuarioPorid).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":5000", router))
}
