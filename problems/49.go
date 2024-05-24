package problems

import "github.com/camnwalter/project-euler/utils"

func FortyNine() int {
	primes := utils.PrimeSieve(10_000)

	var p1 int
	var p2 int
	var p3 int
outer:
	for _, p1 = range primes {
		if p1 == 1487 {
			continue
		}

		for _, p2 = range primes {
			if p2 == p1 || !utils.IsPermutation(p1, p2) {
				continue
			}

			for _, p3 = range primes {
				if p3 == p1 || p3 == p2 || !utils.IsPermutation(p1, p3) {
					continue
				}

				if p1-p2 == p2-p3 {
					break outer
				}
			}
		}
	}

	return p1*100_000_000 + p2*10_000 + p3
}
