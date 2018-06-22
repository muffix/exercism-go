// Package perfect contains tools around perfect numbers
package perfect

import (
	"fmt"
	"math"
)

// Classification is the result of the integer classification
type Classification int

const (
	// ClassificationDeficient is the constant for deficient integers
	ClassificationDeficient Classification = iota
	// ClassificationPerfect is the constant for perfect integers
	ClassificationPerfect Classification = iota
	// ClassificationAbundant is the constant for abundant integers
	ClassificationAbundant Classification = iota
)

// ErrOnlyPositive is the error returned for non-positive input
var ErrOnlyPositive = fmt.Errorf("Can only classify positive integers")

// Classify returns the classification of an integer
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return -1, ErrOnlyPositive
	}

	if n == 1 {
		return ClassificationDeficient, nil
	}

	sum := sumOfFactors(n)

	if sum == n {
		return ClassificationPerfect, nil
	}

	if sum > n {
		return ClassificationAbundant, nil
	}

	return ClassificationDeficient, nil
}

// sumOfFactors returns the sum of all proper factors of an integer
func sumOfFactors(n int64) (sum int64) {
	sum = 1
	var i int64 = 2

	for ; i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i == n/i {
				sum += i
			} else {
				sum += i + n/i
			}
		}
	}
	return
}
