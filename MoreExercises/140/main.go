package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u user) print() {
	toprint := fmt.Sprintf("%s %s is %d yo and sasy: ", u.First, u.Last, u.Age)
	for _, s := range u.Sayings {
		toprint += s + ","
	}
	fmt.Println(toprint)
}

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

type ByAge []user

func (b ByAge) Len() int           { return len(b) }
func (b ByAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByAge) Less(i, j int) bool { return b[i].Age < b[j].Age }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(ByAge(users))

	for _, user := range users {
		user.print()
	}

	sort.Sort(ByLast(users))

	for _, user := range users {
		user.print()
	}
}
