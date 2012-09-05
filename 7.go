package main
import (
	"fmt"
)

func primes(n int, c chan int) {
	primes := make([]int, 1)
	primes[0] = 2

foreach_i:
	for i := 3; ; i += 2 {
		for _,p := range primes {
			if i % p == 0 { // not a prime
				continue foreach_i
			}
		}
		primes = append(primes, i)
		c <- i
		if len(primes) == n {
			break;
		}
	}
	close(c)
}


func main() {
	n, i := 10001, 0
	c := make(chan int)

	go primes(n, c)
	for i = range c {
	}
	fmt.Printf("The %dth prime number is %d.\n", n, i)
}