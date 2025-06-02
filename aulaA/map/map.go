package main

import "fmt"

func main() {
	// --- MAPS ---
	// Maps são coleções de pares chave-valor.
	// As chaves devem ser únicas e de um tipo comparável (string, int, float, bool, etc.).
	// Os valores podem ser de qualquer tipo.
	// map[tipo_da_chave]tipo_do_valor
	fmt.Println("--- MAPS ---")

	// 1. Declaração de um map vazio (valor zero é nil):
	var idades map[string]int // Map de strings para inteiros
	fmt.Println("Map vazio (nil):", idades)

	// 2. Inicialização de um map com 'make':
	// make(map[tipo_da_chave]tipo_do_valor)
	capitais := make(map[string]string) // Map de strings para strings
	fmt.Println("Map de capitais (vazio, com make):", capitais)

	// 3. Adicionando pares chave-valor:
	capitais["Brasil"] = "Brasília"
	capitais["Argentina"] = "Buenos Aires"
	capitais["Chile"] = "Santiago"
	fmt.Println("Capitais após adição:", capitais)

	// 4. Acessando valores por chave:
	fmt.Printf("Capital do Brasil: %s\n", capitais["Brasil"])
	fmt.Printf("Capital do Chile: %s\n", capitais["Chile"])

	// 5. Verificando se uma chave existe (idiomático em Go):
	// valor, ok := map[chave]
	// 'ok' será true se a chave existir, false caso contrário.
	capitalFranca, existe := capitais["França"]
	fmt.Printf("Capital da França: '%s', Existe: %t\n", capitalFranca, existe) // Saída: '', false

	capitaoArgentina, existe := capitais["Argentina"]
	fmt.Printf("Capital da Argentina: '%s', Existe: %t\n", capitaoArgentina, existe) // Saída: 'Buenos Aires', true

	// 6. Inicialização de map com valores literais:
	// map[tipo_da_chave]tipo_do_valor{chave1: valor1, chave2: valor2, ...}
	estoque := map[string]int{
		"Maçãs":    10,
		"Bananas":  25,
		"Laranjas": 15,
	}
	fmt.Println("Estoque:", estoque)

	// 7. Deletando um par chave-valor:
	// delete(map, chave)
	delete(estoque, "Bananas")
	fmt.Println("Estoque após deletar Bananas:", estoque)

	// Tentando acessar uma chave deletada:
	bananasRestantes, existeBananas := estoque["Bananas"]
	fmt.Printf("Bananas restantes: %d, Existe: %t\n", bananasRestantes, existeBananas)

	// 8. Iterando sobre um map com 'range':
	fmt.Println("\nIterando sobre o map de capitais:")
	for pais, capital := range capitais {
		fmt.Printf("O país %s tem a capital %s\n", pais, capital)
	}

	// Iterando apenas sobre as chaves:
	fmt.Println("\nIterando apenas sobre as chaves do estoque:")
	for item := range estoque {
		fmt.Printf("Item em estoque: %s\n", item)
	}

	// Iterando apenas sobre os valores:
	fmt.Println("\nIterando apenas sobre os valores do estoque:")
	for _, quantidade := range estoque { // Usamos '_' para ignorar a chave
		fmt.Printf("Quantidade em estoque: %d\n", quantidade)
	}

	// Tamanho do map:
	fmt.Printf("Número de itens no estoque: %d\n", len(estoque))
}
