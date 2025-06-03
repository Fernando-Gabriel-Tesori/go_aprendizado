package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func envelhecer(p *Pessoa) {
	p.Idade += 1
}

func main() {
	p := Pessoa{"Sky", 25}
	envelhecer(&p)
	fmt.Println(p.Nome, "agora tem", p.Idade, "anos")
}
