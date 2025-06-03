package main

import "fmt"

func main() {
	usuarios := map[string]int{"Sky": 20, "Alice": 30, "Bob": 20}

	for nome, idade := range usuarios {
		fmt.Printf("%s tem %d anos.\n", nome, idade)
	}
}
