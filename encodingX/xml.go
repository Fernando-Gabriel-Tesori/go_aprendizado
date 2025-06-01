package main

import (
	"encoding/xml"
	"fmt"
)

type Cliente struct {
	XMLName xml.Name `xml:"cliente"`
	Nome    string   `xml:"nome"`
	Idade   int      `xml:"idade"`
}

func main() {
	// Struct → XML
	c := Cliente{"", "Maria", 32}
	xmlBytes, _ := xml.MarshalIndent(c, "", "  ")
	fmt.Println("XML:\n", string(xmlBytes))

	// XML → Struct
	xmlTexto := `<cliente><nome>João</nome><idade>45</idade></cliente>`
	var cliente Cliente
	xml.Unmarshal([]byte(xmlTexto), &cliente)
	fmt.Println("Struct:", cliente)
}
