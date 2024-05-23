package problems

import (
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func ThirtyFive() int {
	primes := utils.PrimeSieve(1_000_000)

	count := 0

	for _, prime := range primes {
		if isCircularPrime(prime) {
			count++
		}
	}

	return count
}

func isCircularPrime(n int) bool {
	if !utils.IsPrime(n) {
		return false
	}

	digits := utils.ToDigitArray(n)

	for i := 0; i < len(digits); i++ {
		if !utils.IsPrime(utils.FromDigitArray(digits)) {
			return false
		}

		digit := digits[0]
		digits = slices.Delete(digits, 0, 1)
		digits = append(digits, digit)
	}

	return true
}
