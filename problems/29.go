package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/big"
	"github.com/camnwalter/project-euler/utils"
)

func TwentyNine() {
	min := 2
	max := 100

	uniques := make([]string, 0)

	for a := min; a <= max; a++ {
		aBig := big.New(a)

		for b := min; b <= max; b++ {
			out := big.New(1)
			for range b {
				out = out.Times(aBig)
			}
			uniques = utils.AddIfAbsent(uniques, out.String())
		}
	}

	fmt.Println("uniques =", len(uniques))
}
