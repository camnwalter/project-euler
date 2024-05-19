package problems

import (
	"fmt"
)

func Seven() {
	nthPrime := 0

	prime := 2

	for ; nthPrime < 10_000; nthPrime++ {
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
	if n <= 1 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
