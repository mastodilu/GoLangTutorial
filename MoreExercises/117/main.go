package main

import (
	"fmt"

	"./polygon"
)

func main() {
	var circle polygon.Circle
	circle.Radius = 5
	square := polygon.Square{Side: 5}

	circleInfo, err := polygon.Info(circle)
	if err != nil {
		panic(err)
	}

	squareInfo, err := polygon.Info(square)
	if err != nil {
		panic(err)
	}

	fmt.Println(circleInfo)
	fmt.Println(squareInfo)
}
