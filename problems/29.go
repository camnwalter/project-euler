package problems

import (
	"fmt"
	"math/big"
	"slices"
)

func TwentyNine() {
	uniques := make([]string, 0)

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			a := big.NewInt(int64(a))
			b := big.NewInt(int64(b))

			value := a.Exp(a, b, nil)

			if !slices.Contains(uniques, value.String()) {
				uniques = append(uniques, value.String())
			}
		}
	}

	fmt.Println("uniques =", len(uniques))
}
