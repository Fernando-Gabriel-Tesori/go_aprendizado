package main

import (
	"errors" // Importa o pacote 'errors' para criar novos erros
	"fmt"
)

func main() {
	// --- TRATAMENTO DE ERROS ---
	// Em Go, erros são valores. Funções que podem falhar geralmente retornam
	// um segundo valor do tipo 'error'. Se a operação for bem-sucedida, o erro
	// retornado é 'nil'. Caso contrário, é um valor de erro que descreve a falha.

	fmt.Println("--- TRATAMENTO DE ERROS ---")

	// Exemplo 1: Função que pode retornar um erro
	resultado, err := dividirNumeros(10, 2)
	if err != nil { // É idiomático verificar se o erro NÃO é nil
		fmt.Printf("Erro: %s\n", err)
	} else {
		fmt.Printf("Divisão de 10 por 2: %.2f\n", resultado)
	}

	resultado2, err2 := dividirNumeros(10, 0)
	if err2 != nil {
		fmt.Printf("Erro: %s\n", err2) // Saída: Erro: não pode dividir por zero
	} else {
		fmt.Printf("Divisão de 10 por 0: %.2f\n", resultado2)
	}

	// Exemplo 2: Criando um erro personalizado
	status, erroPersonalizado := verificarIdade(15)
	if erroPersonalizado != nil {
		fmt.Printf("Erro ao verificar idade: %s\n", erroPersonalizado)
	} else {
		fmt.Printf("Status da idade: %s\n", status)
	}

	status2, erroPersonalizado2 := verificarIdade(20)
	if erroPersonalizado2 != nil {
		fmt.Printf("Erro ao verificar idade: %s\n", erroPersonalizado2)
	} else {
		fmt.Printf("Status da idade: %s\n", status2)
	}

	// Exemplo 3: Usando errors.Is e errors.As para comparar erros
	// errors.Is: verifica se um erro na cadeia de erros corresponde a um erro específico.
	// errors.As: desempacota um erro para uma variável de um tipo específico.

	fmt.Println("\n--- Comparando Erros ---")

	// Definição de um tipo de erro personalizado (struct que implementa a interface error)
	type ErroValidacao struct {
		Campo   string
		Mensagem string
	}

	func (e *ErroValidacao) Error() string {
		return fmt.Sprintf("Erro de validação no campo '%s': %s", e.Campo, e.Mensagem)
	}

	// Função que retorna um erro personalizado
	func validarUsuario(nome string, email string) error {
		if nome == "" {
			return &ErroValidacao{Campo: "nome", Mensagem: "Nome não pode ser vazio"}
		}
		if email == "" {
			return &ErroValidacao{Campo: "email", Mensagem: "Email não pode ser vazio"}
		}
		return nil
	}

	errUsuario := validarUsuario("", "teste@exemplo.com")
	if errUsuario != nil {
		fmt.Printf("Erro ao validar usuário: %s\n", errUsuario)

		// Usando errors.Is para verificar se é um ErroValidacao
		var ev *ErroValidacao
		if errors.As(errUsuario, &ev) { // errors.As tenta converter errUsuario para *ErroValidacao
			fmt.Printf("  É um ErroValidacao! Campo: %s, Mensagem: %s\n", ev.Campo, ev.Mensagem)
		}
	}

	errUsuario2 := validarUsuario("Fulano", "")
	if errUsuario2 != nil {
		fmt.Printf("Erro ao validar usuário: %s\n", errUsuario2)
		var ev *ErroValidacao
		if errors.As(errUsuario2, &ev) {
			fmt.Printf("  É um ErroValidacao! Campo: %s, Mensagem: %s\n", ev.Campo, ev.Mensagem)
		}
	}
}

// Função que retorna um float64 e um erro
func dividirNumeros(a, b float64) (float64, error) {
	if b == 0 {
		// errors.New cria um novo erro com a mensagem fornecida.
		return 0, errors.New("não pode dividir por zero")
	}
	return a / b, nil // Retorna o resultado e nil (sem erro)
}

// Função que retorna uma string e um erro
func verificarIdade(idade int) (string, error) {
	if idade < 18 {
		return "", errors.New("idade mínima de 18 anos exigida")
	}
	return "Maior de idade", nil
}