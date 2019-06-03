package main

import (
	"fmt"
)

func yess(msg ...string) {
	fmt.Println(msg)
}

func main() {
	// variadic parameters
	yell("hello")
	yell("Hi", "bye")
	yell("1", "2", "3", "asdafdsfadfaljkdfljk")
}
