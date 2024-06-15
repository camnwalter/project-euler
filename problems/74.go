package problems

import (
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func SeventyFour() int {
	count := 0

	memo := make(map[int]int)

	for n := 1; n < 1_000_000; n++ {
		if digitFactorialChain(n, memo) == 60 {
			count++
		}
	}

	return count
}

func digitFactorialChain(n int, memo map[int]int) int {
	seen := make([]int, 0)

	for !slices.Contains(seen, n) {
		seen = append(seen, n)

		if memo[n] != 0 {
			n = memo[n]
			continue
		}

		sum := 0
		for _, digit := range utils.ToDigitArray(n) {
			sum += utils.Factorial(digit)
		}

		memo[n] = sum
		n = memo[n]
	}

	return len(seen)
}
