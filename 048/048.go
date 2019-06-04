package main

import (
	"fmt"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("BCRYPT: permette di salvare dati in modo criptato")

	password := "qwerty123"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(hashedPassword)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success - password matches")
	}

	fmt.Println("-------------------------------------------")

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte("asdasdasd"))
	if err != nil {
		fmt.Println("You can't log in.")
	} else {
		fmt.Println("you're logged in.")
	}

}
