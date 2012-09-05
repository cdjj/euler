package main

import (
	"fmt"
	"math"
)

// Find all the primes < n, naively.
// This takes about 100s for n=2M.
func primes(n int, c chan int) {
	primes := make([]int, 1)
	primes[0] = 2
	c <- 2

foreach_i:
	for i := 3; i < n; i += 2 {
		for _, p := range primes {
			if i%p == 0 { // not a prime
				continue foreach_i
			}
		}
		primes = append(primes, i)
		c <- i
	}
	close(c)
}

// Find all the primes < n, using Erastotsthenes' sieve.
// This takes about 0.4s for n=2M, but is bounded by RAM.
func fast_primes(n int, c chan int) {
	composite := make([]bool, n)

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if composite[i] { // keep walking until we hit the next prime
			continue
		}

		for j := i * i; j < n; j += i {
			composite[j] = true
		}
	}

	// Deal with 0, 1 as special cases.
	composite[0], composite[1] = true, true

	for i, cmp := range composite {
		if !cmp {
			c <- i
		}
	}
	close(c)
}

func main() {
	n, p := 2000000, 0
	sum := int64(0)

	c := make(chan int)

	go fast_primes(n, c)
	for p = range c {
		sum += int64(p)
	}

	fmt.Printf("The sum of the prime numbers < %d is %d.\n", n, sum)
}
