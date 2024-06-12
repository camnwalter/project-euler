package problems

import (
	"github.com/camnwalter/project-euler/utils"
)

func Seventy() int {
	primes := utils.PrimeSieve(10_000_000)
	minRatio := 1000.
	min := 0

	for n := 2; n < 10_000_000; n++ {
		phi := totient(n, primes)
		ratio := float64(n) / float64(phi)

		if ratio < minRatio {
			if utils.IsPermutation(n, phi) {
				minRatio = ratio
				min = n
			}
		}
	}

	return min
}
