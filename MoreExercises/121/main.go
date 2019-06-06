package main

import (
	"fmt"
)

func doSomething(f1 func(), f2 func(text string) int) string {
	f1()

	len := f2("Hey there!")

	return fmt.Sprint("len is ", len)
}

func main() {
	f1 := func() {
		fmt.Println("I'm a void callback function")
	}

	f2 := func(text string) int {
		return len(text)
	}

	fmt.Println(doSomething(f1, f2))
}
