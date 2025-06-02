package main

import "fmt"

func main() {
	// --- PONTEIROS ---
	// Um ponteiro armazena o endereço de memória de uma variável.
	// Em Go, ponteiros são usados para passar valores por referência,
	// o que significa que a função pode modificar o valor original.

	fmt.Println("--- PONTEIROS ---")

	// 1. Declarando uma variável:
	numero := 10
	fmt.Printf("Valor de 'numero': %d\n", numero)
	fmt.Printf("Endereço de memória de 'numero': %p\n", &numero) // '&' obtém o endereço de memória

	// 2. Declarando um ponteiro:
	// var nome_do_ponteiro *tipo_de_dado
	var ptr *int                                          // ptr é um ponteiro para um inteiro
	fmt.Printf("Valor inicial de 'ptr' (nil): %v\n", ptr) // Ponteiros não inicializados são 'nil'

	// 3. Atribuindo o endereço de uma variável a um ponteiro:
	ptr = &numero
	fmt.Printf("Valor de 'ptr' (endereço de 'numero'): %p\n", ptr)

	// 4. Desreferenciando um ponteiro (acessando o valor que o ponteiro aponta):
	// '*' desreferencia o ponteiro.
	fmt.Printf("Valor apontado por 'ptr': %d\n", *ptr)

	// 5. Modificando o valor através do ponteiro:
	*ptr = 20 // Altera o valor na localização de memória que ptr aponta
	fmt.Printf("Novo valor de 'numero' (modificado via ponteiro): %d\n", numero)
	fmt.Printf("Novo valor apontado por 'ptr': %d\n", *ptr)

	// 6. Ponteiros em funções:
	fmt.Println("\n--- Ponteiros em Funções ---")
	valorOriginal := 5
	fmt.Printf("Antes da função: valorOriginal = %d\n", valorOriginal)

	// Passando o endereço de 'valorOriginal' para a função
	dobrarValor(&valorOriginal)
	fmt.Printf("Depois da função: valorOriginal = %d\n", valorOriginal) // O valor foi alterado!

	// Exemplo de quando não usar ponteiro (passagem por valor):
	valorCopia := 7
	fmt.Printf("Antes da função (cópia): valorCopia = %d\n", valorCopia)
	dobrarCopia(valorCopia) // Passa uma cópia do valor
	fmt.Printf("Depois da função (cópia): valorCopia = %d (não mudou)\n", valorCopia)

	// 7. Ponteiros para structs:
	type Ponto struct {
		X, Y int
	}

	p := Ponto{1, 2}
	ptrP := &p // ptrP é um ponteiro para o struct Ponto

	fmt.Printf("\nStruct original: %+v\n", p) // %+v imprime o struct com nomes dos campos
	fmt.Printf("Ponteiro para struct: %p\n", ptrP)
	fmt.Printf("Valor apontado pelo ponteiro para struct: %+v\n", *ptrP)

	// Acessando campos de um struct através de um ponteiro (sintaxe abreviada):
	// Go permite usar '.' diretamente em ponteiros para structs.
	ptrP.X = 10
	fmt.Printf("Struct após modificar X via ponteiro: %+v\n", p)
}

// Função que recebe um ponteiro para um inteiro.
// Isso permite que a função modifique o valor original da variável.
func dobrarValor(num *int) {
	*num = *num * 2 // Desreferencia 'num' e dobra seu valor
	fmt.Printf("Dentro da função (ponteiro): *num = %d\n", *num)
}

// Função que recebe um inteiro por valor.
// Ela opera em uma CÓPIA do valor original.
func dobrarCopia(num int) {
	num = num * 2 // Altera apenas a cópia local
	fmt.Printf("Dentro da função (cópia): num = %d\n", num)
}
