package problems

import (
	"github.com/camnwalter/project-euler/utils"
)

func SeventyEight() int {
	dp := make([]int, 1)
	dp[0] = 1

	for n := 1; ; n++ {
		dp = append(dp, 0)

		if partitions(n, dp, 1_000_000) == 0 {
			return n
		}
	}
}

func partitions(target int, dp []int, mod int) int {
	if target < 0 {
		return 0
	}

	if dp[target] != 0 {
		return dp[target]
	}

	sum := 0

	for k := 1; k <= target; k++ {
		pentagon := utils.Pentagon(k)
		if pentagon > target {
			break
		}

		a := partitions(target-pentagon, dp, mod)
		if k%2 == 1 {
			sum += a
		} else {
			sum -= a
		}

		pentagon += k
		if pentagon > target {
			break
		}

		b := partitions(target-pentagon, dp, mod)
		if k%2 == 1 {
			sum += b
		} else {
			sum -= b
		}
	}

	dp[target] = sum % mod
	return dp[target]
}
