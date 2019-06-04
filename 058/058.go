package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Error handling.")
	if err != nil {
		fmt.Println("There's an error")
	}
	fmt.Println("Written", n, "bytes")
}
