package main
import (
	"fmt"
)

func square_sum(n int) int {
	i, sum := 0, 0
	for i <= n {
		sum += i
		i++
	}
	return sum * sum
}

func sum_squares(n int) int {
	i, sum := 0, 0

	for i <= n {
		sum += i*i
		i++
	}
	return sum
}

func main() {
	n := 100
	ssq := sum_squares(n)
	sqs := square_sum(n)

	fmt.Printf("The sum of the squares of the first %d natural numbers is %d.\n", n, ssq)
	fmt.Printf("The square of the sum of the first %d natural numbers is %d.\n", n, sqs)
	fmt.Printf("The difference is %d.\n", sqs - ssq)
}