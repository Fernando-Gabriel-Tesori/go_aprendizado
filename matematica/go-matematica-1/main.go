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
	var num1, num2 float64

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)

	fmt.Println("Soma:", somar(num1, num2))
	fmt.Println("Subtração:", subtrair(num1, num2))
	fmt.Println("Multiplicação:", multiplicar(num1, num2))
	fmt.Println("Divisão:", dividir(num1, num2))
}
