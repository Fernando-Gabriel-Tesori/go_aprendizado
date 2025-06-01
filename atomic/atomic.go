package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var contador int64 = 0

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&contador, 1)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementar(&wg)
	}
	wg.Wait()
	fmt.Println("Contador final:", contador)
}
