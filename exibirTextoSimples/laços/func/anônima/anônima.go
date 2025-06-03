package main

import "fmt"

func main() {
	saudacao := func(nome string) {
		fmt.Println("olá", nome, "!")
	}
	saudacao("Sky") //chamando a função anônima
}
