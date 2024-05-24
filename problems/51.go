package problems

import (
	"math"

	"github.com/camnwalter/project-euler/utils"
)

func FiftyOne() int {
	primes := utils.PrimeSieve(1_000_000)
	primeMap := make(map[int]bool)
	for _, prime := range primes {
		primeMap[prime] = true
	}

	var family []int

outer:
	for _, prime := range primes {
		if prime < 1000 {
			continue
		}

		numDigits := 1 + int(math.Log10(float64(prime)))

		for a := 0; a < numDigits; a++ {
			for b := a + 1; b < numDigits; b++ {
				for c := b + 1; c < numDigits; c++ {
					family = make([]int, 0)

					for replacement := 0; replacement <= 9; replacement++ {
						newNum := replaceDigits(prime, replacement, a, b, c)

						if primeMap[newNum] && newNum >= 1000 {
							family = utils.AddIfAbsent(family, newNum)
						}
					}

					if len(family) == 8 {
						break outer
					}
				}
			}
		}
	}

	return utils.Min(family...)
}

func replaceDigits(num int, replacement int, digitA int, digitB int, digitC int) int {
	digits := utils.ToDigitArray(num)
	digits[digitA] = replacement
	digits[digitB] = replacement
	digits[digitC] = replacement

	return utils.FromDigitArray(digits)
}
