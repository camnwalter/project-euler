package problems

import "github.com/camnwalter/project-euler/utils"

func Seven() int {
	nthPrime := 0

	prime := 2

	for ; nthPrime < 10_000; nthPrime++ {
		prime = nextPrime(prime)
	}

	return prime
}

func nextPrime(n int) int {
	var i int
	for i = n + 1; !utils.IsPrime(i); i++ {
	}

	return i
}
