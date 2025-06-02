package main

import (
	"fmt"
	"time"
)

func esperarDados(c chan int) {
	select {
	case msg := <-c:
		fmt.Println("Recido:", msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Tempo limite atingido! Encerrado goroutine.")
	}
}

func main() {
	canal := make(chan int)
	go esperarDados(canal)

	time.Sleep(time.Second * 3) // simula um atraso
}
