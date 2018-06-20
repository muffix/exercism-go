// Package cryptosquare contains tools to encode a message in square code
package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// Encode encodes a message in square code
func Encode(plain string) string {
	plain = strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(plain, ""))

	rows, columns := dimensions(len(plain))

	var encoded strings.Builder

	for c := 0; c < columns; c++ {
		for r := 0; r < rows; r++ {
			index := r*columns + c
			if index < len(plain) {
				encoded.WriteByte(plain[index])
			} else {
				encoded.WriteString(" ")
			}
		}
		if c != columns-1 {
			encoded.WriteString(" ")
		}
	}

	return encoded.String()
}

// dimensions returns the dimensions of the square for a string of a given length
func dimensions(length int) (rows, columns int) {
	sqrt := math.Sqrt(float64(length))
	rows = int(math.Floor(sqrt))
	columns = int(math.Ceil(sqrt))

	if rows*columns < length {
		rows++
	}

	return
}
