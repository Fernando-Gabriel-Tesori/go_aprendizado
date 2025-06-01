package main

import (
	"fmt"
)

func somar(a, b float64) float64 {
	return a + b
}

func subtrair(a, b float64) float64 {
	return a - b
}

func multiplicar(a, b float64) float64 {
	return a * b
}

func dividir(a, b float64) float64 {
	if b == 0 {
		fmt.Println("⚠️ Erro: divisão por zero")
		return 0
	}
	return a / b
}

func main() {
	for {
		var num1, num2 float64
		var opcao int

		fmt.Println("Digite o primeiro número:")
		fmt.Scanln(&num1)

		fmt.Println("Digite o segundo número:")
		fmt.Scanln(&num2)

		fmt.Println("Escolha a operação:")
		fmt.Println("1 - Soma")
		fmt.Println("2 - Subtração")
		fmt.Println("3 - Multiplicação")
		fmt.Println("4 - Divisão")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Printf("Resultado: %.2f\n", somar(num1, num2))
		case 2:
			fmt.Printf("Resultado: %.2f\n", subtrair(num1, num2))
		case 3:
			fmt.Printf("Resultado: %.2f\n", multiplicar(num1, num2))
		case 4:
			fmt.Printf("Resultado: %.2f\n", dividir(num1, num2))
		default:
			fmt.Println("Opção inválida")
		}

		fmt.Println("Deseja realizar outro cálculo? (s/n)")
		var continuar string
		fmt.Scanln(&continuar)
		if continuar != "s" {
			fmt.Println("Encerrando calculadora. Até logo!")
			break
		}
	}
}
