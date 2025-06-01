package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "   Fernando Tesori   "
	nomeLimpo := strings.TrimSpace(strings.ToLower(nome))

	fmt.Println("Nome limpo:", nomeLimpo)
	if strings.Contains(nomeLimpo, "tesori") {
		fmt.Println("O nome contém 'tesori'")
	} else {
		fmt.Println("Não contém 'tesori'")
	}
}
