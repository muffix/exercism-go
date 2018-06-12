// Package diffsquares contains functions around sums and squares
package diffsquares

import "math"

// SquareOfSums returns the square of the sum of the first n integers
func SquareOfSums(n int) int {
	return int(math.Pow(float64(n*(n+1)/2), 2))
}

// SumOfSquares returns the sum of the squares of the first n integers
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns the difference between the square of sums and the sum of squares of the first n integers
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
