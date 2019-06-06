// Package vehicles provides useful structs to work with pseudo vehicle types.
package vehicles

import "fmt"

// Vehicle is the type that represents a vehicle
type Vehicle struct {
	Doors int
	Color string
}

// Truck is the type that represents a truck
type Truck struct {
	Vehicle
	FourWheel bool
}

// Sedan is the type that represents a sedan
type Sedan struct {
	Vehicle
	Luxury bool
}

// String returns a formatted string to print Vehicle type
func (v Vehicle) String() string {
	return fmt.Sprintf("doors: %d, color: %s", v.Doors, v.Color)
}

// String returns a formatted string to print Sedan type
func (s Sedan) String() string {
	return fmt.Sprintf("sedan - %s, luxury: %v", s.Vehicle.String(), s.Luxury)
}

// String returns a formatted string to print Truck type
func (t Truck) String() string {
	return fmt.Sprintf("truck - %s, four wheels: %v", t.Vehicle.String(), t.FourWheel)
}
