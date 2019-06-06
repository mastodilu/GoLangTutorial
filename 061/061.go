package main

/*
	- i test vanno scritti in un file chiamato ..._test.go
	- il file di test va scritto nello stesso package
	- la funzione che testa la funzione Xxx si chiama "func TestXxx(t *testing.T)"

	- per avviare il test si usa "go test"
*/

import (
	"fmt"
)

func MySum(xi ...int) int {
	var sum int
	for _, x := range xi {
		sum += x
	}
	return sum
}

func main() {
	fmt.Println("Test.")

	xi := []int{1, 2, 3, 4}
	fmt.Printf("sum of %v is %d\n", xi, MySum(xi...))
}
