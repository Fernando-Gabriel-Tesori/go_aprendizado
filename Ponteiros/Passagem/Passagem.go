package main

import "fmt"

func dobrar(num *int) {
	*num = *num * 2
}

func main() {
	x := 5
	dobrar(&x)
	fmt.Println("Valor de x após dobrar", x)
}
