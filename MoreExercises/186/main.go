package main

import (
	"fmt"

	"./dog"
)

func main() {
	fmt.Println("0 -", dog.Year(0))
	fmt.Println("1 -", dog.Year(1))
	fmt.Println("5 -", dog.Year(5))
	fmt.Println("10 -", dog.Year(10))
}
