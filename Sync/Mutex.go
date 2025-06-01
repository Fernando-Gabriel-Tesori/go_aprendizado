package main

import (
	"fmt"
	"sync"
)

var (
	saldo = 0
	mu    sync.Mutex
)

func depositar(valor int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	saldo += valor
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go depositar(1, &wg)
	}

	wg.Wait()
	fmt.Println("Saldo final:", saldo)
}
