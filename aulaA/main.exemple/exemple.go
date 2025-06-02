package main // Go é parte de um pacote. "main" é especial para executáveis.

import (
	"fmt" //Import o pacote 'fmt' que fornece funções para formatação de I/0 (entrada/saida).
)

//A func 'main' é o ponto de entrada dequalquer programa Go executável.
//Quando executa seu programa, ele começa aqui.

func main() {
	//fmt.Println é func do pacote 'fmt' que imprime uma linha de texto no console.
	fmt.Println("Olá, Go! Meu primeiro programa")
}
