package main

import "fmt"

func main() {
	var nome string
	var idade int

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome) // leitura do teclado

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	fmt.Printf("Olá, %s! Você tem %d anos.\n", nome, idade)
}
