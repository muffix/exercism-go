// Package summultiples computes the sum of unique multiples of particular numbers
package summultiples

// SumMultiples returns the sum of all numbers upto a given imit that are multiples of the specified numbersl
func SumMultiples(limit int, factors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		if isMultiple(i, factors) {
			sum += i
		}
	}

	return
}

// isMultiple returns whether the candidate is a multiple of any of the given factors
func isMultiple(candidate int, factors []int) bool {
	for _, factor := range factors {
		if candidate%factor == 0 {
			return true
		}
	}
	return false
}
