package main

import "fmt"

func main() {
	i := 1

	var p *int

	p = &i // Pega o endereco da variavel
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}
