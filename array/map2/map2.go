package main

import "fmt"

func main() {
	funcsESalario := map[string]float64{
		"José João":    11325.45,
		"Gabriel":      15456.78,
		"Pedro Junior": 1200.0,
	}

	funcsESalario["Rafael Filho"] = 1350.0 // Adiciona um novo elemento ao map
	delete(funcsESalario, "Inexistente")   // Não da erro ao tentar deletar um valor inexistente

	for nome, salario := range funcsESalario {
		fmt.Println(nome, salario)
	}
}
