// Package prime contains tools to determine prime numbers
package prime

func generator(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func primeFilter(in <-chan int, out chan<- int, p int) int {
	for {
		i := <-in

		if i%p != 0 {
			out <- i
		}
	}
}

// Nth returns the n-th prime number
func Nth(n int) (int, bool) {
	primes := make(chan int)

	go generator(primes)

	p := -1

	for j := 0; j < n; j++ {
		p = <-primes

		notDivisibleByP := make(chan int)

		go primeFilter(primes, notDivisibleByP, p)

		primes = notDivisibleByP
	}

	return p, p > 0
}
