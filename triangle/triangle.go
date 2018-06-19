// Package triangle contains functions around triangles
package triangle

import (
	"math"
)

// Kind is the type that represents the kind of a triangle
type Kind int

const (
	// NaT -  not a triangle
	NaT = iota
	// Equ - equilateral
	Equ = iota
	// Iso - isosceles
	Iso = iota
	// Sca - scalene
	Sca = iota
)

// KindFromSides returns the kind of the triangle from the three sides
func KindFromSides(a, b, c float64) Kind {
	ok := validate(a) && validate(b) && validate(c)
	if !ok || a <= 0 || b <= 0 || c <= 0 || a+b < c || b+c < a || a+c < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}

func validate(a float64) bool {
	return math.IsNaN(a) && math.IsInf(a, 0)
}
