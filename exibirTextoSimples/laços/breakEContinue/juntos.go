package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue //Pula o número 3
		}
		if i == 7 {
			break //Para loop ao chegar no &
		}
		fmt.Println("Número", i)
	}
}
