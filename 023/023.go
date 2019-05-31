package main

import "fmt"

func main() {
	//maps: set di coppie chiave-valore

	//tipo: map[string]int
	// è un composite literal
	//mappa := map[string]int{}

	numbers := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(numbers)

	fmt.Println(numbers["one"])
	fmt.Println(numbers["two"])

	// controlla se la chiave esiste
	v, ok := numbers["three"]
	fmt.Println("three", v, ok)

	v, ok = numbers["two"]
	fmt.Println("two", v, ok)
}
