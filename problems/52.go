package problems

import "github.com/camnwalter/project-euler/utils"

func FiftyTwo() int {
	var x int
	for x = 1; ; x++ {
		good := true
		for n := 2; n <= 6; n++ {
			if !utils.IsPermutation(x, x*n) {
				good = false
				break
			}
		}

		if good {
			break
		}
	}

	return x
}
