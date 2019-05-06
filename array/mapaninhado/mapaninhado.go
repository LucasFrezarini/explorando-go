package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"L": {
			"Lucas Frezarini": 1234.56,
			"Lucao Frezarini": 45545.23,
		},
		"M": {
			"Midoriya Izuku": 5424.52,
		},
		"N": {
			"Naruto": 12345.67,
		},
	}

	delete(funcsPorLetra, "N")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, ":")

		for nome, salario := range funcs {
			fmt.Printf("	%s, salario: %.2f\n", nome, salario)
		}
	}
}
