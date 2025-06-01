package main

import (
	"fmt"
	"math"
)

func potencia(base, expoente float64) float64 {
	return math.Pow(base, expoente)
}

func raizQuadrada(num float64) float64 {
	return math.Sqrt(num)
}

func logNatural(num float64) float64 {
	if num <= 0 {
		fmt.Println("Erro: número deve ser maior que 0")
		return math.NaN()
	}
	return math.Log(num)
}

func media(valores []float64) float64 {
	soma := 0.0
	for _, v := range valores {
		soma += v
	}
	return soma / float64(len(valores))
}

func desvioPadrao(valores []float64) float64 {
	med := media(valores)
	var somaQuadrados float64
	for _, v := range valores {
		somaQuadrados += math.Pow(v-med, 2)
	}
	return math.Sqrt(somaQuadrados / float64(len(valores)))
}

func seno(ang float64) float64 {
	return math.Sin(ang)
}

func cosseno(ang float64) float64 {
	return math.Cos(ang)
}

func tangente(ang float64) float64 {
	return math.Tan(ang)
}

func lerValores() []float64 {
	var n int
	fmt.Print("Quantos valores deseja informar? ")
	fmt.Scanln(&n)

	valores := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Digite o valor #%d: ", i+1)
		fmt.Scanln(&valores[i])
	}
	return valores
}

func main() {
	for {
		fmt.Println("\n📊 CALCULADORA INTERMEDIÁRIA EM GO")
		fmt.Println("1 - Potência")
		fmt.Println("2 - Raiz Quadrada")
		fmt.Println("3 - Logaritmo Natural")
		fmt.Println("4 - Média")
		fmt.Println("5 - Desvio Padrão")
		fmt.Println("6 - Seno")
		fmt.Println("7 - Cosseno")
		fmt.Println("8 - Tangente")
		fmt.Println("9 - Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			var base, expoente float64
			fmt.Print("Base: ")
			fmt.Scanln(&base)
			fmt.Print("Expoente: ")
			fmt.Scanln(&expoente)
			fmt.Println("Resultado:", potencia(base, expoente))
		case 2:
			var num float64
			fmt.Print("Número: ")
			fmt.Scanln(&num)
			fmt.Println("Resultado:", raizQuadrada(num))
		case 3:
			var num float64
			fmt.Print("Número: ")
			fmt.Scanln(&num)
			fmt.Println("Resultado:", logNatural(num))
		case 4:
			valores := lerValores()
			fmt.Printf("Média: %.2f\n", media(valores))
		case 5:
			valores := lerValores()
			fmt.Printf("Desvio Padrão: %.2f\n", desvioPadrao(valores))
		case 6:
			var ang float64
			fmt.Print("Ângulo (em radianos): ")
			fmt.Scanln(&ang)
			fmt.Printf("Seno: %.4f\n", seno(ang))
		case 7:
			var ang float64
			fmt.Print("Ângulo (em radianos): ")
			fmt.Scanln(&ang)
			fmt.Printf("Cosseno: %.4f\n", cosseno(ang))
		case 8:
			var ang float64
			fmt.Print("Ângulo (em radianos): ")
			fmt.Scanln(&ang)
			fmt.Printf("Tangente: %.4f\n", tangente(ang))
		case 9:
			fmt.Println("Encerrando o programa...")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
