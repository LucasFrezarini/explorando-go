package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice != de array! Slice define um pedaço de array
	s2 := a2[1:3] // Pega do index 1 até o 3, NÃO incluindo o ultimo indice (3)
	fmt.Println(a2, s2)

	s3 := a2[:2] // pega da posição inicial até (mas não inclusivo) o index 2
	fmt.Println(a2, s3)

	// Slice de slices! Slice é basicamente definido pelo tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
