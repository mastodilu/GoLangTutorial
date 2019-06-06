// Package dog contains a functions to handle dog years
package dog

// Year converts human years to dog years
func Year(years uint) uint {
	switch years {
	case 0:
		return 0
	default:
		return 1 + (years-1)*7
	}
}
