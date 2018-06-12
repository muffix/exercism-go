// Package hamming contains tools to calculate the Hamming distance
package hamming

import (
	"fmt"
)

// Distance returns the hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("Strings %s and %s are not of equal lenth", a, b)
	}

	distance := 0

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
