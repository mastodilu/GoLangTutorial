package vehicles

import (
	"fmt"
	"testing"
)

func TestVehicleString(t *testing.T) {
	vehicle := Vehicle{Doors: 4, Color: "green"}
	expected := "doors: 4, color: green"
	if got := vehicle.String(); got != expected {
		t.Errorf("Got: '%s'; expected: %s\n", got, expected)
	}
}

func TestTruckString(t *testing.T) {
	truck := Truck{
		Vehicle:   Vehicle{Doors: 4, Color: "green"},
		FourWheel: true,
	}
	expected := "truck - doors: 4, color: green, four wheels: true"
	if got := truck.String(); got != expected {
		t.Errorf("Got: '%s'; expected: %s\n", got, expected)
	}
}

func TestSedanString(t *testing.T) {
	sedan := Sedan{
		Vehicle: Vehicle{Doors: 2, Color: "shocking pink"},
		Luxury:  true,
	}
	expected := "sedan - doors: 2, color: shocking pink, luxury: true"
	if got := sedan.String(); got != expected {
		t.Errorf("Got: '%s'; expected: %s\n", got, expected)
	}
}

func BenchmarkVehicleString(b *testing.B) {
	vehicle := Vehicle{Doors: 2, Color: "shocking pink"}
	for i := 0; i < b.N; i++ {
		vehicle.String()
	}
}

func BenchmarkSedanString(b *testing.B) {
	sedan := Sedan{
		Vehicle: Vehicle{Doors: 2, Color: "shocking pink"},
		Luxury:  true,
	}
	for i := 0; i < b.N; i++ {
		sedan.String()
	}
}

func BenchmarkTruckString(b *testing.B) {
	truck := Truck{
		Vehicle:   Vehicle{Doors: 4, Color: "green"},
		FourWheel: true,
	}
	for i := 0; i < b.N; i++ {
		truck.String()
	}
}

func ExampleVehicleString() {
	vehicle := Vehicle{Doors: 4, Color: "green"}
	fmt.Println(vehicle.String())
	// Output:
	// doors: 4, color: green
}

func ExampleTruckString() {
	truck := Truck{
		Vehicle:   Vehicle{Doors: 4, Color: "green"},
		FourWheel: true,
	}
	fmt.Println(truck.String())
	// Output:
	// truck - doors: 4, color: green, four wheels: true
}

func ExampleSedanString() {
	sedan := Sedan{
		Vehicle: Vehicle{Doors: 2, Color: "shocking pink"},
		Luxury:  true,
	}
	fmt.Println(sedan.String())
	// Output:
	// sedan - doors: 2, color: shocking pink, luxury: true
}
