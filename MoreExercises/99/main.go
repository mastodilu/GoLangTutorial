package main

import (
	"fmt"

	"./vehicles"
)

func main() {
	truck := vehicles.Truck{
		Vehicle:   vehicles.Vehicle{Doors: 4, Color: "dark green"},
		FourWheel: true,
	}
	sedan := vehicles.Sedan{
		Vehicle: vehicles.Vehicle{Doors: 2, Color: "shocking pink"},
		Luxury:  false,
	}
	fmt.Println(sedan.String())
	fmt.Println(truck.String())
	fmt.Println(truck.Vehicle.String())

	fmt.Println("truck.FourWheel", truck.FourWheel)
	fmt.Println("sedan.Luxury", sedan.Luxury)
}
