package main

import "fmt"

func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}

	return "reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
