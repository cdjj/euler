package main
import (
	"fmt"
	"math"
)

func factors(n int64) []int64 {
	max_factor := int64(math.Sqrt(float64(n)))
	factors := make([]int64, 1, max_factor)
	factors[0] = 9223372036854775807 // avoid div-by-zero in first i % factors[f] check.

foreach_i:
	for i := int64(2); i <= max_factor; i++ { 
		if (n % i == 0) { // n is a multiple of i; now check against factors
		for f := range factors {
				if ((i % factors[f]) == 0) { // multiple of a known factor
					continue foreach_i
				}
			}
			// factor is prime
			factors = append(factors, i)
			fmt.Printf("Adding %d to %v\n", i, factors)
		}
	}
	return factors[1:]
}

func main() {
	n := int64(600851475143)
	factors := factors(n)
	fmt.Println("The prime factors of ", n, " are ", factors, 
		" and the largest is ", factors[len(factors) - 1])
}