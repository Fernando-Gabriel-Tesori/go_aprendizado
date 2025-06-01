package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

func main() {
	// Struct → JSON
	p := Produto{"Camiseta", 79.90}
	jsonBytes, _ := json.Marshal(p)
	fmt.Println("JSON:", string(jsonBytes))

	// JSON → Struct
	jsonTexto := `{"nome":"Notebook","preco":3499.99}`
	var produto Produto
	json.Unmarshal([]byte(jsonTexto), &produto)
	fmt.Println("Struct:", produto)
}
