package main

import (
	"fmt"
)

func main() {

	// 💥 func expression
	// le funzioni sono tipi!!
	myfunc := func(x int) {
		fmt.Println("Func expression")
		fmt.Println("The meaning of life:", x)
	}
	myfunc(42)
}
