package main

import "fmt"

/* Funcão executada automáticamente antes da inicialização de cada arquivo .go */
func init() {
	fmt.Println("Inicializando")
}

func main() {
	fmt.Println("Main...")
}
