package main

import "fmt"

func main() {
	x := 1
	y := 2
	// Apenas posfix, nd de --x ou ++x

	x++
	fmt.Println(x)

	y--
	fmt.Println(y)

	// fmt.Println(x == y--) <- n pode
}
