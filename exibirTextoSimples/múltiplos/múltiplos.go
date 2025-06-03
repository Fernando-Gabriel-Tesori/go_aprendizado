package main

import (
	"fmt"
)

func main() {
	dia := "sábado"

	switch dia {
	case "segunda":
		fmt.Println("comeco da semana!")
	case "sábado", "domingo":
		fmt.Println("Fim de semana ❤️‍🔥")
	default:
		fmt.Println("Dia normal.")
	}
}
