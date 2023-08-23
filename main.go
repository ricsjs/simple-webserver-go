package main

import (
	//importação do módulo http
	"encoding/json"
	"net/http"
)

// define a função handler
// w permite escrever a resposta HTTP para o cliente
// r é um ponteiro para uma estrutura que contém informações sobre a requisição HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	book := Book{Title: "Nome do livro", Author: "Autor do livro", Pages: 304}
	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/book", GetBook)
	//especifica qual porta o servidor irá rodar
	http.ListenAndServe(":8080", nil)
}
