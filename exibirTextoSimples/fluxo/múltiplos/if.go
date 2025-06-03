package main

import "fmt"

func main() {
	dia := "sabado"

	switch dia {
	case "segunda":
		fmt.Println("Inicio da semana")
	case "sabado", "domingo":
		fmt.Println("Fim de semana ❤️‍🔥")
	default:
		fmt.Println("Dia normal")
	}
}
