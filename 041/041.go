package main

import (
	"fmt"
)

func main() {
	// 💥 pointers
	a := 42
	b := "ciao"
	fmt.Printf("a:%d type:%T\n", a, a)
	fmt.Printf("&a:%p type:%T\n", &a, &a)
	fmt.Printf("b:%s type:%T\n", b, b)
	fmt.Printf("&b:%p type:%T\n", &b, &b)

	p := &a
	fmt.Printf("p:%v type:%T\n", p, p)
	fmt.Printf("*p:%v type:%T\n", *p, *p)
}
