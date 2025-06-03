package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Soma", soma(1, 2, 3, 4, 5))
}
