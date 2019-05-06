package main

import "fmt"

func main() {
	numeros := [...]int{1, 3, 5, 7, 9} // Compilador vai contar o tamanho do array de acordo com a quantidade de elementos inicializados

	for i, numero := range numeros { // Equivalente ao foreach
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // Ignorando o index
		fmt.Printf("%d ", num)
	}
}
