package main

import "fmt"

func main() {
	names := []string{"matteo", "alberto", "mario"}
	gelati := []string{"cioccolato", "vaniglia", "fragola"}

	multiDimensionalSlice := [][]string{names, gelati}

	fmt.Println(multiDimensionalSlice)
}
