package main

import "fmt"

func main() {
	usuarios := []string{"Sky", "Alice", "Bob", "Eve"}
	idades := []int{25, 30, 19, 40}
	cargos := []string{"Admin", "User", "Guest", "User"}

	fmt.Println("Verificando Usuarios...")

	//loop sobre os usuarios
	for i := 0; i < len(usuarios); i++ {
		nome := usuarios[i]
		idade := idades[i]
		cargo := cargos[i]

		if idade < 18 {
			fmt.Println(nome, "Não tem idade suficiente.")
			continue //Pula pra o proximo usuario
		}

		switch cargo {
		case "Admin":
			fmt.Print(nome, "é um adm.")
		case "User":
			fmt.Print(nome, "é um usuario comum.")
		case "Guest":
			fmt.Println(nome, "tem acesso limitado.")
		default:
			fmt.Println(nome, "tem um cargo desconhecido.")
		}

		if cargo == "Admin" {
			fmt.Println("Adm encontrado, encerrando verificação.")
			break
		}
	}
	fmt.Println("Processo concluido!")
}
