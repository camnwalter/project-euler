package problems

import "github.com/camnwalter/project-euler/utils"

func Ten() int {
	primes := utils.PrimeSieve(2_000_000)
	return utils.SumArray(primes)
}
