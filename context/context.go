package main

import (
	"context"
	"fmt"
	"time"
)

func tarefa(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Tarefa cancelada")
			return
		default:
			fmt.Println("Executando...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go tarefa(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

/*
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
Ideal para requisições externas ou funções que não podem travar indefinidamente.
*/
