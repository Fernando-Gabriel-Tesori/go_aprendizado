package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	senha := "123456"

	// Hash da senha
	hash := sha256.Sum256([]byte(senha))
	hashString := hex.EncodeToString(hash[:])

	// Simula entrada do usuário
	senhaDigitada := "123456"
	hashDigitado := sha256.Sum256([]byte(senhaDigitada))

	if hex.EncodeToString(hashDigitado[:]) == hashString {
		// Gera token aleatório
		token := make([]byte, 16)
		rand.Read(token)
		fmt.Println("Login bem-sucedido. Token:", hex.EncodeToString(token))
	} else {
		fmt.Println("Senha incorreta")
	}
}
