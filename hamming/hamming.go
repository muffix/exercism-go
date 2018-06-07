// Package hamming contains tools to calculate the Hamming distance
package hamming

import (
	"fmt"
	"unicode/utf8"
)

// Distance returns the hamming distance between two strings
func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return -1, fmt.Errorf("Strings %s and %s are not of equal lenth", a, b)
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
