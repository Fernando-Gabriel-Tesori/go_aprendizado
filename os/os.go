package main

import (
	"fmt"
	"os"
)

func main() {
	// Verificar se um arquivo existe
	if _, err := os.Stat("teste.txt"); os.IsNotExist(err) {
		fmt.Println("Arquivo não existe.")
	} else {
		fmt.Println("Arquivo existe.")
	}

	// Criar um diretório
	err := os.Mkdir("meu_diretorio", 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	// Criar arquivo dentro do diretório
	file, err := os.Create("meu_diretorio/arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Escrever no arquivo
	file.WriteString("Conteúdo exemplo.\n")
}
