package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Junta pedaços de caminho com separadores corretos (\ ou /)
	caminho := filepath.Join("usuarios", "fotos", "perfil.png")
	fmt.Println("Caminho:", caminho)

	// Caminho absoluto
	abs, _ := filepath.Abs("teste.txt")
	fmt.Println("Absoluto:", abs)

	// Separar diretório e nome do arquivo
	dir := filepath.Dir(abs)
	base := filepath.Base(abs)
	fmt.Println("Diretório:", dir)
	fmt.Println("Arquivo:", base)

	// Verificar extensão
	ext := filepath.Ext("arquivo.tar.gz")
	fmt.Println("Extensão:", ext) // .gz

	// Caminho relativo
	rel, _ := filepath.Rel("/home/fernando/projetos", "/home/fernando/projetos/site/index.html")
	fmt.Println("Relativo:", rel)
}
