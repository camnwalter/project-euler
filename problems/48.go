package problems

import (
	"fmt"
	"math/big"
)

func FortyEight() {
	sum := big.NewInt(0)
	for n := 1; n <= 1000; n++ {
		bigN := big.NewInt(int64(n))
		sum = sum.Add(sum, bigN.Exp(bigN, bigN, nil))
	}

	out := sum.String()
	fmt.Println(out[len(out)-10:])
}
