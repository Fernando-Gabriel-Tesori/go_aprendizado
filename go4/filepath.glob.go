package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	os.Mkdir("docs", 0755)

	for i := 1; i <= 3; i++ {
		nome := filepath.Join("docs", fmt.Sprintf("arquivo%d.txt", i))
		os.WriteFile(nome, []byte("Conteúdo...\n"), 0644)
	}

	// Listar todos os arquivos .txt
	files, _ := filepath.Glob("docs/*.txt")
	for _, file := range files {
		fmt.Println("Encontrado:", file)
	}
}
