package main

import (
	//importação do módulo http
	"net/http"
)

// define a função handler
// w permite escrever a resposta HTTP para o cliente
// r é um ponteiro para uma estrutura que contém informações sobre a requisição HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func main() {
	http.HandleFunc("/", handler)
	//especifica qual porta o servidor irá rodar
	http.ListenAndServe(":8080", nil)
}
