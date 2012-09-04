package main
import (
	"fmt" 
)

func is_palindrome(n int) bool {
	forward := fmt.Sprint(n)

	// Why is reversing a string so hard in Go?
	i := len(forward)
	runes := make([]rune, i)
	for _,rune := range forward{
		i--
		runes[i] = rune
	}
	reverse := string(runes[i:])

	return (forward == reverse)
}

func main() {
	max_palindrome := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			k := i * j
			if (is_palindrome(k)) {
				if (k > max_palindrome) {
					max_palindrome = k
				}
			}
		}
	}

	fmt.Printf("The largest palindrome is %d\n", max_palindrome)
}