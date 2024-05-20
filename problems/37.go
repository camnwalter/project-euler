package problems

import (
	"fmt"
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func ThirtySeven() {
	count := 0
	sum := 0

	for n := 11; count != 11; n++ {
		if IsTruncatedPrime(n) {
			count++
			sum += n
		}
	}

	fmt.Println("truncated primes sum =", sum)
}

func IsTruncatedPrime(n int) bool {
	if !IsPrime(n) {
		return false
	}

	digits1 := utils.ToDigitArray(n)
	digits2 := utils.ToDigitArray(n)

	for len(digits1) > 0 {
		if !IsPrime(utils.FromDigitArray(digits1)) || !IsPrime(utils.FromDigitArray(digits2)) {
			return false
		}

		digits1 = slices.Delete(digits1, 0, 1)
		digits2 = slices.Delete(digits2, len(digits2)-1, len(digits2))
	}

	return true
}
