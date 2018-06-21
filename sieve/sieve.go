// Package sieve implements the Sieve of Eratosthenes
package sieve

import "math"

// Sieve returns the primes up to the limit
func Sieve(limit int) []int {
	var composite = make(map[int]bool)
	var primes []int

	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(limit)))); i++ {
		for j := 2; j*i <= limit; j++ {
			composite[i*j] = true
		}
	}

	for i := 2; i <= limit; i++ {
		if !composite[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
