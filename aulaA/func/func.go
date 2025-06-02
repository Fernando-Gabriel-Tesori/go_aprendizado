package main

import "fmt"

// --- Definição de Funções ---

// 1. Função sem parâmetros e sem retorno:
func saudacao() {
	fmt.Println("Olá do mundo das funções!")
}

// 2. Função com parâmetros e sem retorno:
// (nome_do_parametro tipo_do_parametro)
func imprimirMensagem(mensagem string) {
	fmt.Println("Mensagem: " + mensagem)
}

// 3. Função com parâmetros e um único retorno:
// (nome_do_parametro tipo_do_parametro) tipo_do_retorno
func somar(a int, b int) int {
	resultado := a + b
	return resultado
}

// 4. Função com múltiplos retornos:
// Go permite que funções retornem múltiplos valores.
// (tipo_do_retorno_1, tipo_do_retorno_2)
func calcularOperacoes(x, y int) (int, int) { // x e y são int, retornamos dois int
	soma := x + y
	produto := x * y
	return soma, produto
}

// 5. Função com retornos nomeados:
// Você pode nomear os valores de retorno, o que os declara como variáveis locais
// e os inicializa com o valor zero de seu tipo. Um 'return' vazio retorna esses valores.
func dividir(dividendo, divisor float64) (quociente float64, erro string) {
	if divisor == 0 {
		erro = "Erro: Divisão por zero!" // Atribui valor à variável de retorno 'erro'
		return                           // Retorna os valores nomeados (quociente será 0.0, erro será a string)
	}
	quociente = dividendo / divisor // Atribui valor à variável de retorno 'quociente'
	return                          // Retorna os valores nomeados
}

func main() {
	// --- Chamada de Funções ---

	saudacao() // Chama a função saudacao

	imprimirMensagem("Aprendendo Go é divertido!") // Chama imprimirMensagem com um argumento

	// Chama somar e armazena o resultado em uma variável
	somaTotal := somar(5, 7)
	fmt.Printf("A soma de 5 e 7 é: %d\n", somaTotal)

	// Chama calcularOperacoes e captura os dois valores de retorno
	s, p := calcularOperacoes(10, 3)
	fmt.Printf("Soma: %d, Produto: %d\n", s, p)

	// Se você quiser ignorar um dos valores de retorno, use o underscore '_'
	_, produtoIgnorado := calcularOperacoes(4, 5)
	fmt.Printf("Produto Ignorado (mas capturado): %d\n", produtoIgnorado)

	// Chama dividir e trata o erro
	q, err := dividir(10.0, 2.0)
	if err != "" { // Verifica se a string de erro não está vazia
		fmt.Println(err)
	} else {
		fmt.Printf("Divisão de 10 por 2: %.2f\n", q)
	}

	q2, err2 := dividir(10.0, 0.0)
	if err2 != "" {
		fmt.Println(err2)
	} else {
		fmt.Printf("Divisão de 10 por 0: %.2f\n", q2)
	}
}
