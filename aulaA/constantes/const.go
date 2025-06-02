package main

import "fmt"

func main() {
	// Constantes são valores que não podem ser alterados após a declaração.
	// São declaradas usando a palavra-chave 'const'.
	// O tipo pode ser explícito ou inferido.

	const PI float64 = 3.14159
	const Saudacao string = "Olá, mundo!"
	const Ano int = 2025

	fmt.Println("Constantes:")
	fmt.Printf("PI: %.5f\n", PI)
	fmt.Printf("Saudação: %s\n", Saudacao)
	fmt.Printf("Ano: %d\n", Ano)

	// Constantes sem tipo explícito (inferido pelo compilador):
	const E = 2.71828    // Go infere float64
	const Verdade = true // Go infere bool

	fmt.Printf("E: %f\n", E)
	fmt.Printf("Verdade: %t\n", Verdade)

	// Exemplo de uso de constante em cálculo:
	const raio = 5.0 // Go infere float64
	area := PI * raio * raio
	fmt.Printf("Área de um círculo com raio %.1f: %.2f\n", raio, area)

	// Erro: Não é possível reatribuir um valor a uma constante.
	// PI = 3.0 // Isso causaria um erro de compilação: cannot assign to PI (untyped float constant)
}
