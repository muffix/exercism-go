// Package collatzconjecture implements the Collatz Conjecture
package collatzconjecture

import "fmt"

// CollatzConjecture returns the number os fteps required to reach 1
func CollatzConjecture(n int) (steps int, err error) {
	if n < 1 {
		return -1, fmt.Errorf("n must be greater than 0, got: %d", n)
	}

	for steps = 0; n > 1; steps++ {
		if n%2 == 0 {
			n >>= 1
		} else {
			n = 3*n + 1
		}
	}
	return
}
