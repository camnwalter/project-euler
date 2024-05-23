package problems

import (
	"math"

	"github.com/camnwalter/project-euler/utils"
)

func FortySix() int {
	primes := utils.PrimeSieve(100_001)

	var n int
	for n = 9; ; n += 2 {
		if utils.IsPrime(n) {
			continue
		}

		found := false
		for _, p := range primes {
			if p > n {
				break
			}

			diff := n - p

			if diff%2 != 0 {
				continue
			}

			square := math.Sqrt(float64(diff / 2))

			if math.Round(square) == square {
				found = true
				break
			}
		}

		if !found {
			break
		}
	}

	return n
}
