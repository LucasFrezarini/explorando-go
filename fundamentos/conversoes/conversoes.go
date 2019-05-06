package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// Operações devem ser feitas com variáveis do mesmo tipo
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Conversao para string na verdade pega o caractere da tabela ASCII do numero. Ex: string(97) = 'a'
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	// Tem dois retornos: o número e o possivel erro (ignorado com _)
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	}
}
