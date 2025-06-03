package main

import "fmt"

func executar(f func(string), nome string) {
	f(nome)
}

func saudacao(nome string) {
	fmt.Println("Bem-vindo,", nome, "!")
}

func main() {
	executar(saudacao, "Sky") //passando a função como argumento
}
