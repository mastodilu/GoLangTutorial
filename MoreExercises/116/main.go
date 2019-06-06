package main

import (
	"fmt"

	"./person"
)

func main() {
	matteo := person.Person{
		Firstname: "Matteo",
		Lastname:  "Dilu",
	}
	fmt.Println(matteo.Speak())
}
