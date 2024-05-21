package problems

import (
	"fmt"
)

func Fifty() {
	primes := PrimeSieve(1_000_000)
	primeMap := make(map[int]bool)
	for _, prime := range primes {
		primeMap[prime] = true
	}

	maxPrime := 0
	maxLen := 0

	for i := 0; i < len(primes); i++ {
		startPrime := primes[i]

		currentSum := startPrime

		for j := i + 1; j < len(primes); j++ {
			prime := primes[j]

			currentSum += prime

			if primeMap[currentSum] && j-i > maxLen {
				maxPrime = currentSum
				maxLen = j - i
			}
		}
	}

	fmt.Println("max prime =", maxPrime, "length", maxLen)
}
