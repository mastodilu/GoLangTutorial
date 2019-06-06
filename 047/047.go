package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByName implements sort.Interface for []Person based on the Name field.
type ByName []Person

func (byname ByName) Len() int           { return len(byname) }
func (byname ByName) Swap(i, j int)      { byname[i], byname[j] = byname[j], byname[i] }
func (byname ByName) Less(i, j int) bool { return byname[i].Name < byname[j].Name }

func main() {
	fmt.Println("Custom sort")

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))

	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
