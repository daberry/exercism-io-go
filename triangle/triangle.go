// Package triangle implements the exercism.io prompt by the same name
package triangle

import "math"

//Kind is a datatype that describes the type of triangle
type Kind int

//Constants to map triangle types
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) (k Kind) {
	sides := []float64{a, b, c}

	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
	}

	if a+b < c || a+c < b || b+c < a {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return
}
