package main

import (
	"fmt"
	"unicode"
)

func main() {
	texto := "GoLang123!@#"

	for _, r := range texto {
		if unicode.IsLetter(r) {
			fmt.Printf("%c é uma letra\n", r)
		} else if unicode.IsDigit(r) {
			fmt.Printf("%c é um número\n", r)
		} else if unicode.IsPunct(r) {
			fmt.Printf("%c é pontuação\n", r)
		}
	}
}
