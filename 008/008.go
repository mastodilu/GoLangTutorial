package main

import "fmt"

const a = 10
const b = "matteo"
const c, d = "di", "lucchio"
const (
	e = 10.90
	f = true
)

func main() {
	fmt.Printf("%v is %T\n", a, a)
	fmt.Printf("%v is %T\n", b, b)
	fmt.Printf("%v is %T\n", c, c)
	fmt.Printf("%v is %T\n", d, d)
	fmt.Printf("%v is %T\n", e, e)
	fmt.Printf("%v is %T\n", f, f)
}
