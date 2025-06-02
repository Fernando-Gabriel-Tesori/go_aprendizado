package main

import (
	"fmt"
	"time"
)

func tarefa(nome string) {
	for i := 0; i < 3; i++ {
		fmt.Println(nome, "rodando...")
		time.Sleep(time.Microsecond * 500)
	}
}

func main() {
	go tarefa("Goroutine 1") // executanto concorrente
	go tarefa("Goroutine 2")

	time.Sleep(time.Second * 2) //esperando execução
	fmt.Println("Programa finalizado!")
}
