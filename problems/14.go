package problems

import (
	"fmt"
)

func Fourteen() {
	longestChain := 0
	longestNumber := 0

	for n := 1; n < 1_000_000; n++ {
		collatz := Collatz(n)
		if collatz > longestChain {
			longestChain = collatz
			longestNumber = n
		}
	}

	fmt.Println("n =", longestNumber)
}

func Collatz(n int) int {
	count := 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}

		count++
	}

	return count
}
