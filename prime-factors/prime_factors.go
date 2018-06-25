// Package prime contains tools to determine prime factors
package prime

// Factors returns the prime factors of n
func Factors(n int64) []int64 {
	factors := []int64{}

	for n%2 == 0 {
		factors = append(factors, 2)
		n >>= 1
	}

	var i int64 = 3
	for ; n > 1; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	return factors
}
