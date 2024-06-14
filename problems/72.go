package problems

import "github.com/camnwalter/project-euler/utils"

func SeventyTwo() int {
	count := 0
	primes := utils.PrimeSieve(1_000_001)

	for n := 2; n <= 1_000_000; n++ {
		count += totient(n, primes)
	}

	return count
}
