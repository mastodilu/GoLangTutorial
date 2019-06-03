package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	licenseToKill bool
	// prs           person
	person // se il nome della struct e il tipo coincidono allora non serve specificarlo due volte
}

func main() {
	jamesBond := secretAgent{
		licenseToKill: true,
		person: person{
			first: "james",
			last:  "bond",
			age:   32,
		},
	}

	fmt.Println(jamesBond.person.first, jamesBond.person.last, "is", jamesBond.person.age)
	if jamesBond.licenseToKill {
		fmt.Println("and is licensed to kill.")
	}
}
