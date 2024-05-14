package problems

import "fmt"

func Seven() {
	nthPrime := 1

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
	for i := 2; i < n; {
		if n%i == 0 {
			return false
		}

		i++
	}

	return true
}
