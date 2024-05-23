package problems

import (
	"github.com/camnwalter/project-euler/big"
	"github.com/camnwalter/project-euler/utils"
)

func Twenty() int {
	return utils.SumArray(BigFactorial(100).Digits)
}

func BigFactorial(n int) *big.MyBigInt {
	out := big.New(n)

	for i := 2; i < n; i++ {
		out = out.Times(big.New(i))
	}

	return out
}
