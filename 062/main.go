package main

import (
	"fmt"

	"./mymath"
)

func main() {
	fmt.Println(mymath.MySum(1, 2, 3))
	fmt.Println(mymath.MySum(10, 2, 3))
	fmt.Println(mymath.MySum(1, 2, 30, -20))
}
