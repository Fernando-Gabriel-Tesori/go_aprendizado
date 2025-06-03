package main

import "fmt"

func contador() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

func main() {
	c := contador()
	fmt.Println(c()) //1
	fmt.Println(c()) //2
	fmt.Println(c()) //3
}
