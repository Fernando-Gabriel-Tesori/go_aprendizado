package main

import "fmt"

func main() {
	// --- STRUCTS ---
	// Structs são tipos de dados compostos que agrupam campos (propriedades) de diferentes tipos.
	// São como classes leves em outras linguagens, mas sem métodos associados diretamente na declaração.
	// São usadas para modelar dados.

	// 1. Definição de um struct:
	// type NomeDoStruct struct {
	//     Campo1 Tipo1
	//     Campo2 Tipo2
	//     ...
	// }
	type Pessoa struct {
		Nome   string
		Idade  int
		Cidade string
		Ativo  bool
	}

	fmt.Println("--- STRUCTS ---")

	// 2. Criando uma instância de um struct:
	// a) Usando a sintaxe de inicialização de campos (recomendado):
	p1 := Pessoa{
		Nome:   "Maria",
		Idade:  28,
		Cidade: "Rio de Janeiro",
		Ativo:  true,
	}
	fmt.Println("Pessoa 1:", p1)
	fmt.Printf("Nome: %s, Idade: %d\n", p1.Nome, p1.Idade)

	// b) Inicialização posicional (menos comum e propenso a erros se a ordem mudar):
	p2 := Pessoa{"João", 35, "São Paulo", false}
	fmt.Println("Pessoa 2:", p2)

	// c) Criando um struct com valores zero (default):
	var p3 Pessoa // Todos os campos terão seus valores zero (string="", int=0, bool=false)
	fmt.Println("Pessoa 3 (valores zero):", p3)

	// 3. Acessando e modificando campos:
	p3.Nome = "Carlos"
	p3.Idade = 40
	p3.Ativo = true
	fmt.Println("Pessoa 3 (modificada):", p3)
	fmt.Printf("Nome da Pessoa 3: %s\n", p3.Nome)

	// 4. Structs aninhados:
	type Endereco struct {
		Rua    string
		Numero int
		CEP    string
	}

	type Cliente struct {
		ID       string
		Pessoa            // Campo aninhado (Pessoa é um campo sem nome, então seus campos são promovidos)
		Endereco Endereco // Campo aninhado nomeado
	}

	cli1 := Cliente{
		ID: "C001",
		Pessoa: Pessoa{
			Nome:   "Fernanda",
			Idade:  22,
			Cidade: "Belo Horizonte",
			Ativo:  true,
		},
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Numero: 123,
			CEP:    "30123-456",
		},
	}
	fmt.Println("\nCliente 1:", cli1)
	// Acessando campos do struct aninhado 'Pessoa' diretamente (promoção de campo):
	fmt.Printf("Nome do Cliente: %s, Idade do Cliente: %d\n", cli1.Nome, cli1.Idade)
	// Acessando campos do struct aninhado 'Endereco' via nome do campo:
	fmt.Printf("Endereço do Cliente: %s, %d, CEP: %s\n", cli1.Endereco.Rua, cli1.Endereco.Numero, cli1.Endereco.CEP)

	// 5. Structs como parâmetros de função:
	imprimirPessoa(p1)
	imprimirPessoa(p2)
}

// Função que recebe um struct como parâmetro
func imprimirPessoa(p Pessoa) {
	fmt.Printf("\n--- Detalhes da Pessoa ---\n")
	fmt.Printf("Nome: %s\n", p.Nome)
	fmt.Printf("Idade: %d\n", p.Idade)
	fmt.Printf("Cidade: %s\n", p.Cidade)
	fmt.Printf("Ativo: %t\n", p.Ativo)
}
