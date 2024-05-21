package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func FortySeven() {
	for i := 4; ; i++ {
		found := true
		for n := 0; n < 4; n++ {
			if len(PrimeFactorize(i+n)) != 4 {
				found = false
				break
			}
		}

		if found {
			fmt.Println("first of the 4 =", i)
			break
		}
	}
}

func PrimeFactorize(n int) []int {
	out := make([]int, 0)

	primes := PrimeSieve(n + 1)

	for i := 0; i < len(primes) && n > 1; i++ {
		prime := primes[i]
		if n%prime == 0 {
			out = utils.AddIfAbsent(out, prime)
			i = -1
			n /= prime
		}
	}

	return out
}
