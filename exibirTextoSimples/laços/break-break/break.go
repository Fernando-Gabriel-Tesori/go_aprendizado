//O comando break faz com que um loop pare imediatamente, ignorando as próximas iterações.

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //interrope o loop quando i for 5
		}
		fmt.Println("Número", i)
	}
}
