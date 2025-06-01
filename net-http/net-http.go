package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bem-vindo à minha API em Go!")
	})

	http.ListenAndServe(":8080", nil)
}
