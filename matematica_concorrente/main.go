package main

import (
	"fmt"
	"log"
	"matematica_concorrente/matrix"
)

func main() {
	a, err := matrix.NewMatrix([][]float64{
		{2, 4},
		{1, 3},
	})
	if err != nil {
		log.Fatal(err)
	}

	b, err := matrix.NewMatrix([][]float64{
		{5, 2},
		{6, 1},
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := a.Multiply(b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("📐 Resultado da multiplicação:")
	for _, row := range result.Data {
		fmt.Println(row)
	}
}
