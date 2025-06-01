package main

import (
	"fmt"
	"path"
)

func main() {
	url := "/api/v1/users/profile"

	fmt.Println("Base:", path.Base(url))                      // profile
	fmt.Println("Dir:", path.Dir(url))                        // /api/v1/users
	fmt.Println("Join:", path.Join("/api", "v2", "produtos")) // /api/v2/produtos
}
