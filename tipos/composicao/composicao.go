package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// Pode ter mais métodos/atributos
}

type bwm7 struct{}

func (b bwm7) ligarTurbo() {
	fmt.Println("Vruuuuum...")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("Fazendo baliza...")
}

func main() {
	var carro esportivoLuxuoso = bwm7{}

	carro.fazerBaliza()
	carro.ligarTurbo()
	fmt.Println(carro)
}
