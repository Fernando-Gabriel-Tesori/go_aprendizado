package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String → Int
	num, err := strconv.Atoi("42")
	if err == nil {
		fmt.Println(num + 8) // 50
	}

	// Int → String
	texto := strconv.Itoa(2025)
	fmt.Println("Ano: " + texto)

	// String → Float
	preco, _ := strconv.ParseFloat("99.90", 64)
	fmt.Println(preco * 1.1) // 109.89

	// Float → String
	precoTexto := strconv.FormatFloat(preco, 'f', 2, 64)
	fmt.Println("Preço formatado: " + precoTexto)
}
