package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	leitor := bufio.NewReader(os.Stdin)

	fmt.Print("Digite sua nota: ")
	nota, _ := leitor.ReadString('\n')

	err := os.WriteFile("nota.txt", []byte(nota), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Nota salva com sucesso!")

	// Ler o que foi salvo
	dados, _ := os.ReadFile("nota.txt")
	fmt.Println("Conteúdo do arquivo:")
	fmt.Println(string(dados))
}
