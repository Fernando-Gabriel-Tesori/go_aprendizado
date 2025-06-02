package main

import (
	"fmt" // Importa o pacote fmt para funções de formatação de I/O
)

func main() {
	// Chama fmt.Println para imprimir strings e um número.
	// fmt.Println retorna dois valores:
	// 1. O número de bytes que foram escritos na saída padrão.
	// 2. Um valor de erro (que será nil se não houver erro).
	// As variáveis numerodebytes e erros são declaradas e inicializadas aqui.
	numerodebytes, erros := fmt.Println("hello word!", "oi galera", 100)

	// Imprime os valores retornados pela chamada anterior de fmt.Println.
	// Isso mostrará quantos bytes foram escritos e se houve algum erro.
	fmt.Println(numerodebytes, erros)
}
