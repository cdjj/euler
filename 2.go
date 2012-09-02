package main
import (
	"fmt"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	sum := 0
	f := fib()

	for n := f() ; n < 4000000; n = f() {
		if (n % 2 == 0) {
			sum += n
		}
	}
 
	fmt.Println("The sum of all even Fibonnaci numbers < 4M is ", sum)
}