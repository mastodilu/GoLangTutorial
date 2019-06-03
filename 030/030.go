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

func sumAll(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	// variadic parameters
	fmt.Println("VARIADIC PARAMTERS ARE SLICES!!\n")
	yell("hello")
	yell("Hi", "bye")
	yell("1", "2", "3", "asdafdsfadfaljkdfljk")

	fmt.Println("sum from 1 to 10 is:", sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
