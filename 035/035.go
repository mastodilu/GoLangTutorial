package main

import (
	"fmt"
)

func main() {
	// 💥 funzione anonima
	func(x int) { // parameters
		fmt.Println("Anonymous function...")
		fmt.Println("The meaning of life:", x)
	}(42) // arguments

}
