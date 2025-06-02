package main

import (
	"fmt"
	"time"
)

func tarefa() {
	for i := 0; i < 5; i++ {
		fmt.Println("executando tarefa", i)
		time.Sleep(time.Millisecond * 500) //simula um tempo de execução
	}
}

func main() {
	go tarefa() //iniciando uma garoutine
	fmt.Println("Enquanto isso, o programa continua...")

	time.Sleep(time.Second * 3) //Espera antes de terminar
}
