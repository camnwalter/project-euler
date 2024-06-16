package problems

import "github.com/camnwalter/project-euler/utils"

func SeventySeven() int {
	primes := utils.PrimeSieve(100)
	for n := 2; ; n++ {
		if waysToMake(n, primes) > 5000 {
			return n
		}
	}
}
