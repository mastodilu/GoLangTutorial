package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Printf("%T\n", x)  // tipo di variabile
	fmt.Printf("%b\n", x)  // binario
	fmt.Printf("%x\n", x)  // esadecimale
	fmt.Printf("%#x\n", x) // 0xEsadecimale
	fmt.Printf("%v\n", x)  // valore normale

	/*
		%v	the value in a default format
		when printing structs, the plus flag (%+v) adds field names
	*/
}
