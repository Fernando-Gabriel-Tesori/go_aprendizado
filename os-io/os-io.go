package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Criar um arquivo e escrever algo
	err := os.WriteFile("arquivo.txt", []byte("Olá, Go!\nAprendendo I/O!"), 0644)
	if err != nil {
		panic(err)
	}

	// Abrir e ler conteúdo
	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer arquivo.Close()

	// Ler tudo de uma vez
	conteudo, err := io.ReadAll(arquivo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conteúdo do arquivo:")
	fmt.Println(string(conteudo))
}
