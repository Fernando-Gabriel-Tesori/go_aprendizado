package main

import "fmt"

func main() {
	// 1. Declaração explícita de variável:
	// var nome_da_variavel tipo_de_dado = valor
	var nome string = "Alice"
	var idade int = 30
	var altura float64 = 1.75
	var estaAtivo bool = true

	fmt.Println("Variáveis declaradas explicitamente:")
	fmt.Printf("Nome: %s, Idade: %d, Altura: %.2f, Ativo: %t\n", nome, idade, altura, estaAtivo)

	// 2. Declaração de variável com inferência de tipo (operador curto :=):
	// Go infere o tipo da variável com base no valor atribuído.
	// Só pode ser usado dentro de funções.
	cidade := "São Paulo"
	populacao := 12000000
	temperatura := 25.5

	fmt.Println("\nVariáveis com inferência de tipo:")
	fmt.Printf("Cidade: %s, População: %d, Temperatura: %.1f\n", cidade, populacao, temperatura)

	// 3. Reatribuição de valor (tipo não pode ser alterado):
	nome = "Bob"
	idade = 25
	fmt.Println("\nVariáveis reatribuídas:")
	fmt.Printf("Novo Nome: %s, Nova Idade: %d\n", nome, idade)

	// 4. Declaração de múltiplas variáveis:
	var x, y int = 10, 20
	fmt.Printf("\nMultiplas variáveis: x = %d, y = %d\n", x, y)

	// 5. Declaração de variáveis sem valor inicial (elas terão o 'valor zero' do tipo):
	var padraoInt int       // 0
	var padraoFloat float64 // 0.0
	var padraoString string // "" (string vazia)
	var padraoBool bool     // false

	fmt.Println("\nValores zero (default):")
	fmt.Printf("Int: %d, Float: %.1f, String: '%s', Bool: %t\n", padraoInt, padraoFloat, padraoString, padraoBool)
}
