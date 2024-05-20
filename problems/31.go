package problems

import (
	"fmt"
)

func ThirtyOne() {

	// 200a + 100b + 50c + 20d + 10e + 5f + 2g + h = 200
	target := 200
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}

	numWays := waysToMake(target, coins)

	fmt.Println("number of ways to make 2 pounds =", numWays)
}

func waysToMake(target int, coins []int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= target; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[target]
}
