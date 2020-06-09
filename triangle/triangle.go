// Package triangle implements the exercism.io prompt by the same name
package triangle

import "math"

//Kind is a datatype that describes the type of triangle
type Kind string

//Constants to map triangle types
const (
	NaT = "NaT"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) (k Kind) {
	if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	} else if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		k = NaT
	} else if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
	} else if a+b < c || a+c < b || b+c < a {
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
