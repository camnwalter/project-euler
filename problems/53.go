package problems

import "github.com/camnwalter/project-euler/utils"

func FiftyThree() int {
	count := 0

	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			comb := utils.Combination(n, r)
			if comb > 1_000_000 || comb <= 0 {
				count++
			}
		}
	}

	return count
}
