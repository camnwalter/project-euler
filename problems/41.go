package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func FortyOne() {
	for n := 9; n >= 4; n-- {
		nums := make([]int, 0)
		for i := 1; i <= n; i++ {
			nums = append(nums, i)
		}

		permutations := utils.Permutations(nums)

		primes := utils.Filter(utils.Map(permutations, func(val []int, _ int) int {
			return utils.FromDigitArray(val)
		}), func(val int, _ int) bool {
			return IsPrime(val)
		})

		if len(primes) > 0 {
			fmt.Println("max prime =", utils.Max(primes...))
			break
		}
	}
}
