package main

import (
	"fmt"
	"time" // Para usar funções de tempo, como Sleep
)

// --- CONCORRÊNCIA EM GO ---
// Go foi construído com concorrência em mente, tornando-a muito fácil de usar.
// Os dois conceitos principais são:
// 1. Goroutines: Funções que são executadas concorrentemente. São leves e eficientes.
// 2. Channels: Canais são a forma de goroutines se comunicarem entre si,
//    evitando o compartilhamento de memória diretamente (o que pode levar a problemas).
//    A filosofia de Go é "Não se comunique compartilhando memória; em vez disso,
//    compartilhe memória comunicando-se."

// Função que será executada como uma goroutine
func worker(id int) {
	fmt.Printf("Worker %d: Começando...\n", id)
	time.Sleep(time.Second) // Simula algum trabalho demorado
	fmt.Printf("Worker %d: Terminou.\n", id)
}

// Função que envia mensagens para um canal
func remetente(canal chan<- string, mensagem string) {
	fmt.Printf("Enviando mensagem: %s\n", mensagem)
	canal <- mensagem // Envia a mensagem para o canal
}

// Função que recebe mensagens de um canal
func receptor(canal <-chan string) {
	msg := <-canal // Recebe a mensagem do canal
	fmt.Printf("Mensagem recebida: %s\n", msg)
}

func main() {
	fmt.Println("--- GOROUTINES ---")

	// 1. Iniciando Goroutines:
	// Use a palavra-chave 'go' antes de uma chamada de função para executá-la como uma goroutine.
	// A função 'main' também é uma goroutine.
	go worker(1) // Inicia 'worker(1)' em uma nova goroutine
	go worker(2) // Inicia 'worker(2)' em outra nova goroutine

	// Como as goroutines são executadas em segundo plano, a função 'main' pode terminar
	// antes que elas completem. Usamos Sleep aqui para dar tempo às goroutines.
	// Em aplicações reais, você usaria mecanismos como 'sync.WaitGroup' ou channels
	// para coordenar goroutines.
	time.Sleep(2 * time.Second) // Espera um pouco para as goroutines terminarem
	fmt.Println("Goroutines concluídas (ou main terminou).")

	fmt.Println("\n--- CHANNELS ---")

	// 2. Criando Canais:
	// make(chan tipo_de_dado)
	// Canais são tipados, ou seja, só podem enviar/receber dados de um tipo específico.
	canalMensagens := make(chan string) // Cria um canal que transporta strings

	// 3. Enviando e Recebendo de Canais:
	// 'canal <- valor' : Envia 'valor' para 'canal'
	// 'variavel := <-canal' : Recebe um valor de 'canal' e atribui a 'variavel'
	// Canais são bloqueantes por padrão:
	// - Uma operação de envio bloqueia até que alguém esteja pronto para receber.
	// - Uma operação de recebimento bloqueia até que alguém esteja pronto para enviar.

	// Iniciando goroutines para enviar e receber mensagens
	go remetente(canalMensagens, "Olá do canal!")
	go receptor(canalMensagens)

	time.Sleep(time.Second) // Dá tempo para as goroutines de canal se comunicarem

	// 4. Canais com Buffer:
	// make(chan tipo_de_dado, capacidade_do_buffer)
	// Canais com buffer não bloqueiam imediatamente se o buffer não estiver cheio.
	// Eles bloqueiam apenas quando o buffer está cheio (para envios) ou vazio (para recebimentos).
	canalBuffer := make(chan int, 2) // Canal de inteiros com buffer de capacidade 2

	fmt.Println("\n--- CHANNELS COM BUFFER ---")
	canalBuffer <- 1 // Envia 1 para o canal (não bloqueia, buffer tem espaço)
	canalBuffer <- 2 // Envia 2 para o canal (não bloqueia, buffer tem espaço)
	// canalBuffer <- 3 // Esta linha BLOQUEARIA, pois o buffer está cheio (capacidade 2)

	fmt.Println("Valores no buffer:", <-canalBuffer, <-canalBuffer) // Recebe os valores
	// fmt.Println("Valor restante:", <-canalBuffer) // Esta linha BLOQUEARIA, pois o buffer está vazio

	// Exemplo de uso de canal para sinalizar conclusão:
	done := make(chan bool, 1) // Canal com buffer para sinalizar que uma goroutine terminou

	go func() { // Função anônima como goroutine
		fmt.Println("\nGoroutine de sinalização: Trabalhando...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Goroutine de sinalização: Terminou.")
		done <- true // Envia um sinal para o canal 'done'
	}()

	<-done // Bloqueia até receber um valor do canal 'done'
	fmt.Println("Main: Goroutine de sinalização terminou.")
}
