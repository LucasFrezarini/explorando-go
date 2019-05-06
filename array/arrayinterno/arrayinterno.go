package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 42
	fmt.Println(s1, s2) // A pos 0 dos dois slices vão ser 42, pois apontam para o mesmo array interno
}
