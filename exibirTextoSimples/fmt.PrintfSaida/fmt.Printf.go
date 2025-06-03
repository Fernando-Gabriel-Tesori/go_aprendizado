//fmt.Printf(). Ele permite especificar o formato dos dados, como números e strings.

package main

import "fmt"

func main() {
	nome := "sky"
	idade := 25
	altura := 1.75

	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d anos\n", idade)
	fmt.Printf("Altura: %2.f metros\n", altura)
}
