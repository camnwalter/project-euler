package problems

import "github.com/camnwalter/project-euler/utils"

func FortyFive() int {
	var out int
	for n := 286; ; n++ {
		out = utils.Triangle(n)
		if utils.IsPentagonal(out) && utils.IsHexagonal(out) {
			break
		}
	}

	return out
}
