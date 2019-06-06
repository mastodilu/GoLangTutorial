package polygon

import (
	"fmt"
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 5}
	expected := 78.53981633974483
	if got := circle.Area(); got != expected {
		t.Errorf("\nExpected: '%v'\nGot: '%v'\n", expected, got)
	}
}

func TestCirclePerimeter(t *testing.T) {
	circle := Circle{Radius: 5}
	expected := 31.41592653589793
	if got := circle.Perimeter(); got != expected {
		t.Errorf("\nExpected: '%v'\nGot: '%v'\n", expected, got)
	}
}

func ExampleCircleArea() {
	circle := Circle{Radius: 5}
	fmt.Println(circle.Area())
	// Output:
	// 78.53981633974483
}

func ExampleCirclePerimeter() {
	circle := Circle{Radius: 5}
	fmt.Println(circle.Perimeter())
	// Output:
	// 31.41592653589793
}

// -----------------------------

func TestSquareArea(t *testing.T) {
	square := Square{Side: 5}
	expected := 25.0
	if got := square.Area(); almostEqual(got, expected) {
		t.Errorf("\nExpected: '%v'\nGot: '%v'\n", expected, got)
	}
}

func TestSquarePerimeter(t *testing.T) {
	square := Square{Side: 5}
	expected := 20.0
	if got := square.Perimeter(); almostEqual(got, expected) {
		t.Errorf("\nExpected: '%v'\nGot: '%v'\n", expected, got)
	}
}

func ExampleSquareArea() {
	square := Square{Side: 5}
	fmt.Println(square.Area())
	// Output:
	// 25
}

func ExampleSquarePerimeter() {
	square := Square{Side: 5}
	fmt.Println(square.Perimeter())
	// Output:
	// 20
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) == math.SmallestNonzeroFloat64
}

// -----------------------------

func TestInfo(t *testing.T) {
	circle := Circle{Radius: 5}
	expected := "Circle(5)-Area:78.53981633974483,Perimeter:31.41592653589793"
	if got, _ := Info(circle); got != expected {
		t.Errorf("\nExpected: '%v'\nGot: '%v'\n", expected, got)
	}
}

func ExampleInfo() {
	square := Square{Side: 5}

	squareInfo, _ := Info(square)

	fmt.Println(squareInfo)
	// Output:
	// Square(5)-Area:25,Perimeter:20
}

// ----------------------------

func BenchmarkInfo(b *testing.B) {
	square := Square{Side: 5}
	for i := 0; i < b.N; i++ {
		Info(square)
	}
}

func BenchmarkArea(b *testing.B) {
	square := Square{Side: 5}
	for i := 0; i < b.N; i++ {
		square.Area()
	}
}

func BenchmarkPerimeter(b *testing.B) {
	circle := Circle{Radius: 5}
	for i := 0; i < b.N; i++ {
		circle.Area()
	}
}
