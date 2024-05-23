package problems

import (
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func TwentyOne() int {
	tracked := make([]int, 0)

	for a := 1; a < 10_000; a++ {
		b := d(a)

		if a != b && b >= 1 && b < 10_000 && d(b) == a && !slices.Contains(tracked, a) && !slices.Contains(tracked, b) {
			tracked = append(tracked, a, b)
		}
	}

	return utils.SumArray(tracked)
}

func d(n int) int {
	return utils.SumArray(utils.GetProperDivisors(n))
}
