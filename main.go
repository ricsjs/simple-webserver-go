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

	books := []Book{ // Cria uma slice de livros
		{Title: "Nome do livro", Author: "Autor do livro", Pages: 304},
		{Title: "Outro livro", Author: "Outro autor", Pages: 256},
	}

	books = append(books, Book{Title: "Mais um livro", Author: "Mais um autor", Pages: 312}) // Adiciona mais um livro à slice

	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/book", GetBook)
	//especifica qual porta o servidor irá rodar
	http.ListenAndServe(":8080", nil)
}
