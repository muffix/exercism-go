// Package allyourbase implements a base converter
package allyourbase

import (
	"fmt"
	"math"
)

// ConvertToBase converts a number from one base to another
func ConvertToBase(inBase int, digits []int, outBase int) ([]int, error) {
	if inBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}

	if outBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}

	if len(digits) == 0 {
		return []int{0}, nil
	}

	var dec int
	for i := 0; i <= len(digits)-1; i++ {
		digit := digits[i]
		if digit >= inBase || digit < 0 {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		dec += digit * intPow(inBase, len(digits)-1-i)
	}

	return convertDec(dec, outBase), nil

}

func intPow(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}

func convertDec(dec, base int) []int {
	if dec == 0 {
		return []int{0}
	}

	var digits []int

	for dec > 0 {
		digits = append([]int{dec % base}, digits...)
		dec /= base
	}

	return digits
}
