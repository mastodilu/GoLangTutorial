package main

import (
	"fmt"
)

func factorialNonRecursive(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// 💥 recursion
	fmt.Println("factorial 10:", factorial(10))
	fmt.Println("factorial 10:", factorialNonRecursive(10))
}
