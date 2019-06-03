package main

import (
	"fmt"
)

// struct, complex data type, aggregate data type
type person struct {
	first string
	last string
	age int
}

func main() {
	p1 := person{
		first: "james",
		last: "bond",
		age: 27,
	}

	p2 := person{
		first: "matteo",
		last: "dilu",
		age: 12,
	}

	fmt.Println(p1)
	fmt.Println(p2.first, p2.last, "is", p2.age)
}
