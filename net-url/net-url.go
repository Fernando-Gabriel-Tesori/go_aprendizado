package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Montar uma URL
	valores := url.Values{}
	valores.Set("busca", "produtos")
	valores.Add("filtro", "ativos")

	u := url.URL{
		Scheme:   "https",
		Host:     "api.meusite.com",
		Path:     "/v1/lista",
		RawQuery: valores.Encode(),
	}

	fmt.Println("URL gerada:", u.String())

	// Parsear URL
	analise, _ := url.Parse("https://site.com/busca?q=go+lang&ordem=asc")
	fmt.Println("Host:", analise.Host)
	fmt.Println("Query:", analise.Query().Get("q")) // go lang
	fmt.Println("Query completa:", analise.RawQuery)
}
