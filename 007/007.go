package main

import (
	"fmt"
)

func main() {
	// ho una stringa
	s := "H"
	fmt.Println(s)

	// la stringa diventa slice di bytes
	bs := []byte(s)
	fmt.Println(bs)

	// salvo il primo byte e lo stampo
	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	// ho un numero in base 10
	num := 667
	fmt.Printf("%T\n", num)  // tipo
	fmt.Printf("%b\n", num)  // binario
	fmt.Printf("%x\n", num)  // esadecimale
	fmt.Printf("%#X\n", num) // esadecimale
}
