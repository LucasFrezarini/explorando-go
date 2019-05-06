package main

import "fmt"

func main() {
	s := make([]int, 10) // Cria um slice de inteiros de 10 posicões
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        // Cria um slice de 10 elementos, apontando para um array de 20 elementos
	fmt.Println(s, len(s), cap(s)) // len = tamanho do slice, cap = capacidade do array que o slice está apontando

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // O slice aumenta de tamanho automaticamente ao atingir seu limite e o array que ele está apontando dobra de tamanho
	fmt.Println(s, len(s), cap(s))
}
