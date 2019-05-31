package main

import (
	"fmt"
)

func main() {

	// SWITCH

	// non viene specificata una condizione. Ogni case è sempre vero, quindi
	// entra in un case casuale tra tutti
	switch {
	case true:
		fmt.Println("a")
	case true:
		fmt.Println("b")
	case true:
		fmt.Println("c")
	default:
		fmt.Println("default")
	}

	fmt.Println("------------------------------")

	// switch classico. Entra dove è vera la condizione o nel case default
	x := 1
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

	fmt.Println("------------------------------")

	// ❗ switch con fallthrough. Non si interrompe alla fine del case in cui entra ma alla fine del successivo.
	y := 1
	switch y {
	case 1:
		fmt.Println("1 fallthrough")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

	fmt.Println("------------------------------")

	// ❗ switch con case multipli.
	z := 1
	switch z {
	case 1, 2, 3:
		fmt.Println("1 or 2 or 3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("default")
	}

}
