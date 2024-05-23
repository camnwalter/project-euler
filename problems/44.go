package problems

import "github.com/camnwalter/project-euler/utils"

func FortyFour() int {
	var difference int

	for i := 1; i < 10_000; i++ {
		for j := i + 1; j < 10_000; j++ {
			p_i := i * (3*i - 1) / 2
			p_j := j * (3*j - 1) / 2

			sum := p_i + p_j
			diff := p_j - p_i

			if utils.IsPentagonal(sum) && utils.IsPentagonal(diff) {
				difference = diff
				break
			}
		}
	}

	return difference
}
