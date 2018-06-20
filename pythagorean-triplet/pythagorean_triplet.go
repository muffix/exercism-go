// Package pythagorean contains functions around pythagorean triplets
package pythagorean

import "math"

// Triplet defines a triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var triplets []Triplet

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			c := int(math.Sqrt(float64(a*a + b*b)))
			if c*c == a*a+b*b && c <= max && a >= min && b >= min && c >= min {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}

	return triplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	candidates := Range(1, p)
	var triplets []Triplet

	for _, t := range candidates {
		if t[0]+t[1]+t[2] == p {
			triplets = append(triplets, t)
		}
	}

	return triplets
}
