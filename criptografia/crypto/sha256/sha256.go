package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	texto := "senha123"
	hash := sha256.Sum256([]byte(texto))
	fmt.Printf("SHA256: %x\n", hash)
}
