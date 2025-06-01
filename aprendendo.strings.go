package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "Fernando Tesori"

	// ToUpper e ToLower
	fmt.Println(strings.ToUpper(nome)) // "FERNANDO TESORI"
	fmt.Println(strings.ToLower(nome)) // "fernando tesori"

	// Contains (verifica se existe um trecho)
	fmt.Println(strings.Contains(nome, "Tesori")) // true

	// Replace
	frase := "banana banana banana"
	fmt.Println(strings.Replace(frase, "banana", "maçã", 2)) // "maçã maçã banana"

	// Split (quebra por espaço)
	palavras := strings.Split(nome, " ")
	fmt.Println(palavras) // ["Fernando", "Tesori"]

	// Join (junta tudo com algo)
	junto := strings.Join(palavras, "-")
	fmt.Println(junto) // "Fernando-Tesori"

	// Trim (remove espaços antes e depois)
	comEspacos := "   Olá mundo   "
	fmt.Println(strings.TrimSpace(comEspacos)) // "Olá mundo"
}

//run code 🫡
//string

//Pacote strconv: texto ↔ número
