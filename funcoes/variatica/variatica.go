package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	if len(numeros) == 0 {
		return 0.0
	}

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(8.5, 7.1, 9.2, 9.9, 7.4))
}
