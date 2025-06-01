package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Criar e escrever
	file, _ := os.Create("produtos.csv")
	writer := csv.NewWriter(file)
	writer.Write([]string{"Nome", "Preço"})
	writer.Write([]string{"Teclado", "129.90"})
	writer.Write([]string{"Mouse", "59.90"})
	writer.Flush()
	file.Close()

	// Ler
	file, _ = os.Open("produtos.csv")
	reader := csv.NewReader(file)
	linhas, _ := reader.ReadAll()
	for i, linha := range linhas {
		fmt.Printf("Linha %d: %v\n", i, linha)
	}
}
