package problems

import (
	"fmt"
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func ThirtyFive() {
	primes := PrimeSieve(1_000_000)

	count := 0

	for _, prime := range primes {
		if IsCircularPrime(prime) {
			count++
		}
	}

	fmt.Println("circular primes =", count)
}

func IsCircularPrime(n int) bool {
	if !IsPrime(n) {
		return false
	}

	digits := utils.ToDigitArray(n)

	for i := 0; i < len(digits); i++ {
		if !IsPrime(utils.FromDigitArray(digits)) {
			return false
		}

		digit := digits[0]
		digits = slices.Delete(digits, 0, 1)
		digits = append(digits, digit)
	}

	return true
}

func PrimeSieve(upperBound int) []int {
	primes := utils.DefaultSlice(upperBound, true)
	primes[0] = false
	primes[1] = false

	for i := 2; i*i < upperBound; i++ {
		if primes[i] {
			for j := i * i; j < upperBound; j += i {
				primes[j] = false
			}
		}
	}

	out := make([]int, 0)
	for num, isPrime := range primes {
		if isPrime {
			out = append(out, num)
		}
	}

	return out
}
