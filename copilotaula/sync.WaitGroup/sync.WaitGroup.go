package main

import (
	"fmt"
	"sync"
)

func tarefa(id int, wg *sync.WaitGroup) {
	defer wg.Done() // marca esta goroutine como concluida
	fmt.Println("Executando tarefa", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; 1 <= 5; i++ {
		wg.Add(1) //adicionar uma goroutine ao grupo
		go tarefa(i, &wg)
	}

	wg.Wait() //Aguarda todas as goroutines finalizarem
	fmt.Println("Todas as tarefas concluidas!")
}
