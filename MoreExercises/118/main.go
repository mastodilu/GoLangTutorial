package main

import (
	"fmt"
)

func whatsup(x int) {
	for i := 0; i < x; i++ {
		fmt.Println("whatsuppp")
	}
}

func main() {

	defer func(msg string) {
		for i := 0; i < 5; i++ {
			fmt.Println(msg)
		}
	}("Good bye")

	fmt.Println("Good morning sunshine")

	f := whatsup
	f(3)
}
