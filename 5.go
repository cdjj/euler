package main
import (
	"fmt"
	"math"
)

/* Return the factors (with multiplicity) for a number. */
func factors(n int64) map[int64]int {
	factors := make(map[int64]int, 1)
	
	for i := int64(2); i <= n; {
		if (n % i == 0) {
			factors[i] += 1
			n = n / i
		} else {
			i++
		}
	}

	return factors
}

func main() {
	n, sum := int64(20), int64(1)
	lcm_factors := make(map[int64]int)

	for i := int64(2); i <= n; i++ {
		factors := factors(i)
		for f,m := range factors {
			lcm_factors[f] = int(math.Max(float64(m), float64(lcm_factors[f])))
		}
	}

	for f,m := range lcm_factors {
		sum *= int64(math.Pow(float64(f), float64(m)))
	}

	fmt.Printf("The smallest number evenly divisible by [1,%d] is %d.\n", n, sum)
}