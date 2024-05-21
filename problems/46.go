package problems

import (
	"fmt"
	"math"
)

func FortySix() {
	primes := PrimeSieve(100_001)

	for n := 9; ; n += 2 {
		if IsPrime(n) {
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
			fmt.Println("can't make", n)
			break
		}
	}
}
