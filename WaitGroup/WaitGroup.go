package main

import (
	"fmt"
	"sync"
	"time"
)

func tarefa(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Iniciando tarefa", id)
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando tarefa", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go tarefa(i, &wg)
	}

	wg.Wait()
	fmt.Println("Todas as tarefas terminaram ✅")
}
