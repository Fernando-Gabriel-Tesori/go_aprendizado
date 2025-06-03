package main

import "fmt"

func dividir(a int, b int) (int, int) {
	quociente := a / b
	resto := a % b
	return quociente, resto
}

func main() {
	q, r := dividir(10, 3)
	fmt.Println("Quaciente:", q, "Resto", r)
}
