package main

import (
	"fmt"
	"math"
)

// Potência: base elevado ao expoente
func potencia(base, expoente float64) float64 {
	return math.Pow(base, expoente)
}

// Raiz quadrada
func raizQuadrada(num float64) float64 {
	return math.Sqrt(num)
}

// Logaritmo natural (base e)
func logaritmoNatural(num float64) float64 {
	if num <= 0 {
		fmt.Println("Erro: logaritmo de número não positivo")
		return math.NaN()
	}
	return math.Log(num)
}

// Média aritmética
func media(valores []float64) float64 {
	soma := 0.0
	for _, v := range valores {
		soma += v
	}
	return soma / float64(len(valores))
}

// Desvio padrão (populacional)
func desvioPadrao(valores []float64) float64 {
	med := media(valores)
	var somaQuadrados float64
	for _, v := range valores {
		somaQuadrados += math.Pow(v-med, 2)
	}
	variancia := somaQuadrados / float64(len(valores))
	return math.Sqrt(variancia)
}

// Funções trigonométricas
func seno(ang float64) float64 {
	return math.Sin(ang)
}

func cosseno(ang float64) float64 {
	return math.Cos(ang)
}

func tangente(ang float64) float64 {
	return math.Tan(ang)
}

// Função principal
func main() {
	valores := []float64{2, 4, 4, 4, 5, 5, 7, 9}

	fmt.Println("Potência 2^3:", potencia(2, 3))
	fmt.Println("Raiz quadrada de 16:", raizQuadrada(16))
	fmt.Println("Logaritmo natural de e:", logaritmoNatural(math.E))

	fmt.Printf("Média dos valores: %.2f\n", media(valores))
	fmt.Printf("Desvio padrão dos valores: %.2f\n", desvioPadrao(valores))

	angulo := math.Pi / 6 // 30 graus em radianos
	fmt.Printf("Seno de 30°: %.2f\n", seno(angulo))
	fmt.Printf("Cosseno de 30°: %.2f\n", cosseno(angulo))
	fmt.Printf("Tangente de 30°: %.2f\n", tangente(angulo))
}
