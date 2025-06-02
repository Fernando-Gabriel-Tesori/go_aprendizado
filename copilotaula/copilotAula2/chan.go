package main

import "fmt"

func enviarDados(c chan string) {
	c <- "mensagem enviada por goroutine!"
}

func main() {
	canal := make(chan string) //Criando um Canal
	go enviarDados(canal)      //Goroutine envia mensagem

	mensagem := <-canal //recebe a mensagem do canal
	fmt.Println(mensagem)
}
