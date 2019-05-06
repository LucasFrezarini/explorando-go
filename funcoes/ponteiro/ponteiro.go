package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n)        // por valor
	fmt.Println(n) // Não é afetado!

	inc2(&n)       // por referência
	fmt.Println(n) // É afetado!
}
