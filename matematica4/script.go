package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// ----------------------------
// MATRIZ: DETERMINANTE (2x2 ou 3x3)
// ----------------------------

func determinante2x2(m [2][2]float64) float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

func determinante3x3(m [3][3]float64) float64 {
	return m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) -
		m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) +
		m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])
}

// ----------------------------
// INTEGRAÇÃO NUMÉRICA: MÉTODO DOS TRAPÉZIOS
// ----------------------------

func integrarTrapz(f func(x float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	soma := 0.5*f(a) + 0.5*f(b)
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		soma += f(x)
	}
	return soma * h
}

// ----------------------------
// DERIVADA NUMÉRICA (Diferença central)
// ----------------------------

func derivadaCentral(f func(x float64) float64, x float64, h float64) float64 {
	return (f(x+h) - f(x-h)) / (2 * h)
}

// ----------------------------
// OPERAÇÕES COM NÚMEROS COMPLEXOS
// ----------------------------

func complexOperations() {
	c1 := complex(3, 4) // 3 + 4i
	c2 := complex(1, -2)

	fmt.Println("\n🔢 Operações com números complexos:")
	fmt.Println("Soma:", c1+c2)
	fmt.Println("Produto:", c1*c2)
	fmt.Println("Módulo de c1:", cmplx.Abs(c1))
	fmt.Println("Fase de c1 (rad):", cmplx.Phase(c1))
}

// ----------------------------
// FUNÇÃO PRINCIPAL
// ----------------------------

func main() {
	fmt.Println("📌 Determinante 2x2:")
	m2 := [2][2]float64{{4, 6}, {3, 8}}
	fmt.Println("Resultado:", determinante2x2(m2))

	fmt.Println("\n📌 Determinante 3x3:")
	m3 := [3][3]float64{
		{6, 1, 1},
		{4, -2, 5},
		{2, 8, 7},
	}
	fmt.Println("Resultado:", determinante3x3(m3))

	fmt.Println("\n📌 Integração numérica de f(x) = x² de 0 a 2:")
	integral := integrarTrapz(func(x float64) float64 {
		return x * x
	}, 0, 2, 100)
	fmt.Printf("Resultado (aprox.): %.5f\n", integral)

	fmt.Println("\n📌 Derivada de f(x) = sin(x) em x = π/4:")
	derivada := derivadaCentral(math.Sin, math.Pi/4, 0.0001)
	fmt.Printf("Resultado (aprox.): %.5f\n", derivada)

	complexOperations()
}
