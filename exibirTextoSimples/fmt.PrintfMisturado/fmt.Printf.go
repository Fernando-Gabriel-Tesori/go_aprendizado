package main

import "fmt"

func main() {
	nome := "sky"
	idade := 25
	altura := 1.75
	ativo := true

	fmt.Println("Dados do usuario:")
	fmt.Printf("Nome: %s, idade: %d, altura: %.2f, ativo: %t\n", nome, idade, altura, ativo)
}
