package main

import (
	"fmt"
	"sync"
)

var contador int
var mutex sync.Mutex

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock() //Bloqueia o acesso
	contador++
	mutex.Unlock() //Libera o acesso
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementar(&wg)
	}

	wg.Wait()
	fmt.Print("Contador final:", contador)
}
