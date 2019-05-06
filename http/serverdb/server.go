package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	fmt.Println("Executando...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
