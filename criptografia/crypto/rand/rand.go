package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	bytes := make([]byte, 16) // 128 bits
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	token := hex.EncodeToString(bytes)
	fmt.Println("Token seguro:", token)
}
