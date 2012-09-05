package main

import (
	"fmt"
)

func main() {

	// A straightforward yet n^3 algorithm; it should only take us a few hundred million
	// candidates to find our triple.  Fortunately, CPUs are fast.

	for c := 1000; c > 0; c-- {
		for b := c - 1; b > 0; b-- {
			for a := b - 1; a > 0; a-- {
				if (a*a + b*b == c*c) && (a + b + c == 1000) {
					fmt.Printf("[a, b, c] = [%d, %d, %d]; a*b*c = %d\n",
						a, b, c, a * b * c)
					return
				}
			}
		}
	}
}