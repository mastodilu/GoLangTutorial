// Package mymath contains useful mathematical functions.
package mymath

// MySum calculates the sum over a set of numbers.
func MySum(xi ...int) int {
	var sum int
	for _, x := range xi {
		sum += x
	}
	return sum
}
