// Package grains contains functions to calculate the number of rice grains on a chess board
package grains

import "fmt"

// Square returns the number of grains on a given square number
func Square(field int) (uint64, error) {
	if field < 1 || field > 64 {
		return 0, fmt.Errorf("Field number must be between 1 and 64. Got: %d", field)
	}

	return 1 << uint(field-1), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	return 1<<64 - 1
}
