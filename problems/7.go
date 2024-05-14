package problems

import (
	"fmt"
	"math"
)

func Seven() {
	nthPrime := 0

	prime := 2

	for ; nthPrime < 10_001; nthPrime++ {
		prime = NextPrime(prime)
	}

	fmt.Println("nth prime =", prime)
}

func NextPrime(n int) int {
	i := n + 1
	for {
		if IsPrime(i) {
			break
		}

		i++
	}

	return i
}

func IsPrime(n int) bool {
	if n == 2 || n == 3 || n == 5 || n == 7 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i <= int(math.Floor(math.Sqrt(float64(n)))); {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}

		i += 6
	}

	return true
}
