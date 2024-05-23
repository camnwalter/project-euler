package problems

import "github.com/camnwalter/project-euler/utils"

func FortySeven() int {
	primes := utils.PrimeSieve(1_000_000)

	var i int
	for i = 4; ; i++ {
		found := true
		for n := 0; n < 4; n++ {
			if len(primeFactorize(i+n, primes)) != 4 {
				found = false
				break
			}
		}

		if found {
			break
		}
	}

	return i
}

func primeFactorize(n int, primes []int) []int {
	out := make([]int, 0)

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
