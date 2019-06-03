package main

import (
	"fmt"
)

// riceve uno slice di string
func yell(msg ...string) {
	fmt.Printf("number of parameters: %d\n", len(msg))
	for i, value := range msg {
		fmt.Printf("i:%d - %s\n", i, value)
	}
}

func main() {
	// variadic parameters
	fmt.Println("VARIADIC PARAMTERS ARE SLICES!!\n")
	yell("hello")
	yell("Hi", "bye")
	yell("1", "2", "3", "asdafdsfadfaljkdfljk")
}
