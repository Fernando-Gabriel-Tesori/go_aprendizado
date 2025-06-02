package main

import "fmt"

func main() {
	// --- ARRAYS ---
	// Arrays em Go têm tamanho fixo e são declarados com o tipo e o tamanho.
	// [tamanho]tipo_de_dado
	fmt.Println("--- ARRAYS ---")

	// Declaração de um array de 5 inteiros. Valores padrão são 0.
	var numeros [5]int
	fmt.Println("Array vazio (valores zero):", numeros) // Saída: [0 0 0 0 0]

	// Atribuindo valores a elementos específicos
	numeros[0] = 10
	numeros[2] = 30
	fmt.Println("Array com valores atribuídos:", numeros) // Saída: [10 0 30 0 0]

	// Acessando elementos
	fmt.Printf("Primeiro elemento: %d\n", numeros[0])
	fmt.Printf("Terceiro elemento: %d\n", numeros[2])

	// Inicialização de array com valores:
	frutas := [3]string{"Maçã", "Banana", "Laranja"}
	fmt.Println("Array de frutas:", frutas)

	// Omitir o tamanho e deixar o compilador contar:
	cores := [...]string{"Vermelho", "Verde", "Azul", "Amarelo"}
	fmt.Println("Array de cores:", cores)
	fmt.Printf("Tamanho do array de cores: %d\n", len(cores)) // len() retorna o tamanho

	// Iterando sobre um array:
	fmt.Println("Iterando sobre o array de frutas:")
	for i, fruta := range frutas {
		fmt.Printf("Índice %d: %s\n", i, fruta)
	}

	// --- SLICES ---
	// Slices são mais poderosos e flexíveis que arrays.
	// Eles são construídos sobre arrays e podem crescer/encolher.
	// Eles são a forma mais comum de usar coleções em Go.
	// []tipo_de_dado (sem tamanho fixo)
	fmt.Println("\n--- SLICES ---")

	// 1. Declaração de um slice vazio (valor zero é nil):
	var meuSlice []int
	fmt.Println("Slice vazio (nil):", meuSlice, "Tamanho:", len(meuSlice), "Capacidade:", cap(meuSlice))

	// 2. Criando um slice com make (tipo, tamanho, capacidade):
	// make([]tipo, tamanho, capacidade)
	// tamanho: número de elementos que o slice contém inicialmente.
	// capacidade: tamanho do array subjacente.
	numerosSlice := make([]int, 3, 5) // Slice de int, com 3 elementos, capacidade de 5
	fmt.Println("Slice com make:", numerosSlice, "Tamanho:", len(numerosSlice), "Capacidade:", cap(numerosSlice))
	numerosSlice[0] = 1
	numerosSlice[1] = 2
	fmt.Println("Slice com valores:", numerosSlice)

	// 3. Inicialização de slice com valores literais:
	// Go infere o tamanho e a capacidade.
	nomes := []string{"Ana", "João", "Maria"}
	fmt.Println("Slice de nomes:", nomes, "Tamanho:", len(nomes), "Capacidade:", cap(nomes))

	// 4. Adicionando elementos a um slice com 'append':
	// 'append' retorna um novo slice (pode ser o mesmo subjacente ou um novo array).
	nomes = append(nomes, "Pedro")
	fmt.Println("Slice de nomes após append:", nomes, "Tamanho:", len(nomes), "Capacidade:", cap(nomes))

	nomes = append(nomes, "Carla", "Lucas") // Adicionar múltiplos elementos
	fmt.Println("Slice de nomes após múltiplos append:", nomes, "Tamanho:", len(nomes), "Capacidade:", cap(nomes))

	// 5. Fatiando um slice (slice de um slice ou array):
	// slice[low:high] -> inclui 'low', exclui 'high'
	// slice[:high] -> do início até 'high' (excluindo)
	// slice[low:] -> de 'low' (incluindo) até o final
	subSlice := nomes[1:4] // Elementos do índice 1 (João) ao 3 (Pedro)
	fmt.Println("Sub-slice de nomes (1:4):", subSlice)

	primeirosDois := nomes[:2]
	fmt.Println("Primeiros dois nomes:", primeirosDois)

	ultimosTres := nomes[len(nomes)-3:]
	fmt.Println("Últimos três nomes:", ultimosTres)

	// 6. Copiando slices:
	// copy(destino, origem)
	copiaNomes := make([]string, len(nomes))
	copy(copiaNomes, nomes)
	fmt.Println("Cópia de nomes:", copiaNomes)

	// Alterar a cópia não afeta o original (se o array subjacente for diferente)
	copiaNomes[0] = "Roberto"
	fmt.Println("Nomes original:", nomes)
	fmt.Println("Cópia alterada:", copiaNomes)
}
