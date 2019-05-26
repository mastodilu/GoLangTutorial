package main

import (
	"fmt"
)

// variabile globale
var a = 10
var b int

func main() {

	c := 20 // variabile locale

	//il tipo viene assegnato in fase di compilazione, non serve specificarlo

	fmt.Println(a + c)

	//variabile ridichiarata perchè è prioritaria da quella dichiarata fuori dal main
	b := "hey"
	fmt.Println(b)
}
