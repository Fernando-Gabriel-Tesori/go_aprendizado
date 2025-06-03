package main

import "fmt"

func main() {

	nota := 90

	if nota >= 90 {
		fmt.Println("Excelente!")
	} else if nota >= 70 {
		fmt.Println("Aprovado!")
	} else {
		fmt.Println("Reprovado.")
	}

}
