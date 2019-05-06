package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string) // Cria um map com chave de inteiros e valores do tipo string

	aprovados[12345678978] = "Lucao"
	aprovados[42345675423] = "Ana Julia"
	aprovados[51236536346] = "Rocket"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[42345675423])
	delete(aprovados, 42345675423) // Remove a chave 42345675423
	fmt.Println(aprovados)
	fmt.Println(aprovados[42345675423])
}
