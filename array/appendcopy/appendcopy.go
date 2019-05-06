package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) # n funciona

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)

	/* Copia o conteúdo do slice1 para o slice2. Não aumenta o tamanho do slice2, ou seja, só vai copiar os 2 primeiros elementos do slice1
	para o slice 2 */
	copy(slice2, slice1)
	fmt.Println(slice2)
}
