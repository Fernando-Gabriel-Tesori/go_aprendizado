package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x

	fmt.Println("Valor de x:", x)
	fmt.Println("Endereço de x:", &x)
	fmt.Println("Valor armazenado no ponteiro p:", *p)
}
