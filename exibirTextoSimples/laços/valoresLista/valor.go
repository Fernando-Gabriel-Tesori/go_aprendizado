package main

import "fmt"

func main() {
	nomes := []string{"Sky", "Alice", "Bob"}

	for indice, nome := range nomes {
		fmt.Printf("Índice: %d, Nome: %s\n", indice, nome)
	}
}
