package main

import (
	"fmt"
)

func main() {
	// 💥 struct anonime
	jamesBond := struct { // dichiaro la struct anonima
		first string
		last  string
		age   int
	}{ // inizializzo la struct anonima
		first: "james",
		last:  "bond",
		age:   32,
	}
	fmt.Println(jamesBond)
}
