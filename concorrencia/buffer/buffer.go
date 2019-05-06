package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	for i := 1; i <= 6; i++ {
		fmt.Printf("Enviando o valor %d para o channel...\n", i)
		ch <- i
		fmt.Printf("Valor %d enviado!\n", i)
	}
}

func main() {
	ch := make(chan int, 3) // Cria um buffer de tamanho 3. A operação de escrita só vai ser bloqueada quando atingir 3 elementos esperando para serem lidos
	go rotina(ch)

	time.Sleep(time.Second * 2)
	fmt.Println(<-ch)
}
