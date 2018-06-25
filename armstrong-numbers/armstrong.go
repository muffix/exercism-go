// Package armstrong contains tools to determine whether a number is an Armstrong number
package armstrong

import "math"

// digits returns the number of digits in n
func digits(n int) int {
	i := 1

	if n >= 1e8 {
		i += 8
		n /= 1e8
	}
	if n >= 1e4 {
		i += 4
		n /= 1e4
	}
	if n >= 1e2 {
		i += 2
		n /= 1e2
	}
	if n >= 10 {
		i++
	}

	return i
}

// IsNumber returns whether n is an Armstrong number
func IsNumber(n int) bool {
	numDigits := digits(n)

	sum := 0
	target := n

	for n > 0 {
		sum += int(math.Pow(float64(n%10), float64(numDigits)))
		n /= 10
	}

	return sum == target
}
