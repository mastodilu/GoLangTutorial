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
	licenseToKill bool // campi dichiarati esplicitamente
	person             // 💥 campi dichiarati implicitamente: se il nome della struct e il tipo coincidono allora non serve specificarlo due volte
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

	// 💥 esempio di "promoted fields"
	// accedo ai campi della struct interna come se appartenessero alla struct esterna
	// ma funziona solo così, non in fase di creazione della variabile
	/* errore:
	matteo := secretAgent{
		licenseToKill: false,
		first: ... //ERRORE
	}
	*/
	fmt.Print(jamesBond.first, " ", jamesBond.last, " is ", jamesBond.age)
	if jamesBond.licenseToKill {
		fmt.Print(" and licensed to kill")
	}
	fmt.Println(".")
}
