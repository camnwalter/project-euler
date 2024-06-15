package problems

import (
	"math"

	"github.com/camnwalter/project-euler/utils"
)

func SeventyThree() int {
	count := 0

	for d := 4.; d <= 12_000; d++ {
		for n := math.Floor(d/3) + 1; n < d/2; n++ {
			if utils.GCF(int(d), int(n)) == 1 {
				count++
			}
		}
	}

	return count
}
