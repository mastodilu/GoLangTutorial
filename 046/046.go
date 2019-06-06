package main

import (
	"fmt"
	"sort"
)

func main() {
	randomNumbers := []int{4, 1, 2, 6, 3, 4, 5, 5, 6, 5, 7, 8, -1}
	randomStrings := []string{"abc", "ABC", "MATTEO", "john"}

	fmt.Println("Sorting")

	fmt.Println(randomNumbers)
	fmt.Println(randomStrings)

	sort.Ints(randomNumbers)
	fmt.Println(randomNumbers)
	sort.Strings(randomStrings)
	fmt.Println(randomStrings)
}
