package main

import "fmt"

func main() {
	saudacao := func(nome string) string {
		return "Olá, " + nome + "!"
	}
	fmt.Println(saudacao("Sky"))
}
