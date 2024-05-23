package problems

import "github.com/camnwalter/project-euler/utils"

func FortyOne() int {
	var primes []int

	for n := 9; n >= 4; n-- {
		nums := make([]int, 0)
		for i := 1; i <= n; i++ {
			nums = append(nums, i)
		}

		permutations := utils.Permutations(nums)

		primes = utils.Filter(utils.Map(permutations, func(val []int, _ int) int {
			return utils.FromDigitArray(val)
		}), func(val int, _ int) bool {
			return utils.IsPrime(val)
		})

		if len(primes) > 0 {
			break
		}
	}

	return utils.Max(primes...)
}
