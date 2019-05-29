package main

import (
	"fmt"
)

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	// for loop annidati
	for i, stop := 0, false; i < 11 && !stop; i++ {
		fmt.Print(i, "-")
		for j := 0; j < 10; j++ {
			fmt.Print(j)
			if i == 9 && j == 9 {
				stop = true
			}
		}
		fmt.Println()
	}
}
