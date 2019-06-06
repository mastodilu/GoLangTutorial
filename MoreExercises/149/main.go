package main

import (
	"fmt"
)

type Person struct {
	First, Last string
	Age         int
}

type Human interface {
	SaySomethingSmart() string
}

func (p *Person) SaySomethingSmart() string {
	return "meh."
}

func say(h Human) {
	fmt.Println(h.SaySomethingSmart())
}

func main() {
	h := Person{
		First: "Matteo",
		Last:  "Dilu",
		Age:   25,
	}
	say(&h) // 💥 usa un pointer receiver e funziona lo stesso
}
