package problems

import (
	"github.com/camnwalter/project-euler/big"
	"github.com/camnwalter/project-euler/utils"
)

func FiftySix() int {
	maxSum := 0

	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			base := big.New(1)

			for range b {
				base = base.Times(big.New(a))
			}

			sum := utils.SumArray(base.Digits)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}
