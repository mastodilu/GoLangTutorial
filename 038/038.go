package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func reviewOldStuff() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(numbers...))
}

// numeri pari
func even(f func(xi ...int) int, vi ...int) int {
	yi := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func main() {
	reviewOldStuff()

	evenNumbers := func(min, max int) func() []int {
		return func() []int {
			even := []int{}

			for i := min; i <= max; i++ {
				if i%2 == 0 {
					even = append(even, i)
				}
			}
			return even
		}

	}

	total := sum(evenNumbers(1, 5)()...)
	fmt.Println(total)

	// --------------------------------------

	// 💥 callback
	// una callback è una funzione che viene passata come parametro di un'altra funzione
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total = sum(numbers...)
	s := even(sum, numbers...)
	fmt.Println(s)
}
