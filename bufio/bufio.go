package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	leitor := bufio.NewReader(os.Stdin)

	fmt.Print("Digite uma frase: ")
	frase, _ := leitor.ReadString('\n') // lê até a quebra de linha

	fmt.Println("Você digitou:", frase)
}
