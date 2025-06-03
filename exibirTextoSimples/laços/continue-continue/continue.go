//O comando continue faz com que a execução pule para a próxima iteração do loop, sem executar o código abaixo dele.

package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue //Pula a interação quando i for 2
		}
		fmt.Println("Número", i)
	}
}
