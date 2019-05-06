package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(i imprimivel) {
	fmt.Println(i.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Dio", "Brando"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça para uma jovem de 16 anos", 300.99}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Outra calça", 301.00}
	imprimir(p2)

}
