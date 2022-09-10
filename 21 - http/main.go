package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina de usuario"))
}

func main() {

	http.HandleFunc("/home", home)

	http.HandleFunc("/user", user)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
