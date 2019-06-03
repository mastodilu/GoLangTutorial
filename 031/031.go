package main

import (
	"fmt"
)

func sumAll(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// variadic parameters
	fmt.Println("sum from 1 to 10 is:", sumAll(numbers...))
}
