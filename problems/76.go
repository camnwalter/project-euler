package problems

func SeventySix() int {
	coins := make([]int, 99)
	for i := range coins {
		coins[i] = i + 1
	}
	return waysToMake(100, coins)
}
