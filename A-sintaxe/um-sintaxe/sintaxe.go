package main //Define o pacote principal

import "fmt" //exibi mensagem

//package main -> indica que o codigo é executavel.
//import "fmt" -> Adiciona funções úteis como Println
//func main() -> ponto de entrada do programa
func idade() {
	var nome string = "sky"   // texto
	var idade int = 20        // idade
	var altura float64 = 1.75 // altura
	var ativo bool = true     // verdadeiro ou falso

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Ativo:", ativo)
}

func main() {
	var idade int = 20

	if idade >= 18 {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é de menor idade")
	}
}
