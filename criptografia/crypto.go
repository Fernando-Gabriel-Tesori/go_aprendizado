package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	texto := "senha123"
	hash := md5.Sum([]byte(texto))
	fmt.Printf("MD5: %x\n", hash) // ⚠️ MD5 é fraco e não recomendado para segurança real
}
