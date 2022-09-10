package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {

	u := user{
		Name:  "John Doe",
		Email: "juan@gmail.com",
	}

	templates.ExecuteTemplate(w, "index.html", u)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
