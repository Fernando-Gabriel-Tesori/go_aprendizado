package main

import "fmt"

func main() {
	// --- IF / ELSE IF / ELSE ---
	fmt.Println("--- IF / ELSE IF / ELSE ---")
	idade := 18

	if idade >= 18 {
		fmt.Println("É maior de idade.")
	} else {
		fmt.Println("É menor de idade.")
	}

	temperatura := 28

	if temperatura > 30 {
		fmt.Println("Está muito quente!")
	} else if temperatura > 20 {
		fmt.Println("A temperatura está agradável.")
	} else {
		fmt.Println("Está frio.")
	}

	// Declaração curta dentro do if (variável 'num' só existe dentro do if/else):
	if num := 10; num%2 == 0 {
		fmt.Printf("%d é um número par.\n", num)
	} else {
		fmt.Printf("%d é um número ímpar.\n", num)
	}

	// --- FOR LOOPS ---
	fmt.Println("\n--- FOR LOOPS ---")

	// 1. For tradicional (como em C/Java):
	fmt.Println("Contagem de 0 a 4:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 2. For como 'while' (condição única):
	fmt.Println("Contagem regressiva de 3:")
	j := 3
	for j > 0 {
		fmt.Println(j)
		j--
	}

	// 3. For infinito (com break):
	fmt.Println("Loop infinito com break:")
	k := 0
	for {
		fmt.Println(k)
		k++
		if k >= 3 {
			break // Sai do loop
		}
	}

	// 4. For com 'range' (para iterar sobre coleções como arrays, slices, maps, strings):
	fmt.Println("Iterando sobre uma string:")
	fruta := "banana"
	for indice, caractere := range fruta {
		fmt.Printf("Índice: %d, Caractere: %c\n", indice, caractere)
	}

	// --- SWITCH CASE ---
	fmt.Println("\n--- SWITCH CASE ---")
	diaDaSemana := "Terça"

	switch diaDaSemana {
	case "Segunda":
		fmt.Println("Início da semana.")
	case "Terça", "Quarta", "Quinta": // Múltiplos valores para um case
		fmt.Println("Meio da semana.")
	case "Sexta":
		fmt.Println("Quase fim de semana!")
	default: // Caso nenhum dos anteriores corresponda
		fmt.Println("Fim de semana ou dia inválido.")
	}

	// Switch sem expressão (avalia condições booleanas):
	pontuacao := 85
	switch {
	case pontuacao >= 90:
		fmt.Println("Nota A")
	case pontuacao >= 80:
		fmt.Println("Nota B")
	case pontuacao >= 70:
		fmt.Println("Nota C")
	default:
		fmt.Println("Nota D ou F")
	}

	// Switch com fallthrough (executa o próximo case também):
	fmt.Println("Exemplo de fallthrough:")
	numero := 2
	switch numero {
	case 1:
		fmt.Println("Case 1")
		fallthrough // Continua para o próximo case
	case 2:
		fmt.Println("Case 2")
		fallthrough // Continua para o próximo case
	case 3:
		fmt.Println("Case 3")
	default:
		fmt.Println("Default")
	}
}
