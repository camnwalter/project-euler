package problems

import (
	"slices"

	"github.com/camnwalter/project-euler/big"
	"github.com/camnwalter/project-euler/utils"
)

func FiftyFive() int {
	count := 0

	for n := 1; n < 10_000; n++ {
		if isLychrel(n) {
			count++
		}
	}

	return count
}

func isLychrel(n int) bool {
	bigN := big.New(n)

	for range 50 {
		bDigits := slices.Clone(bigN.Digits)
		slices.Reverse(bDigits)

		b := &big.MyBigInt{Digits: bDigits, Signum: big.Positive}

		c, _ := bigN.Plus(b)

		if utils.IsPalindrome(c.String()) {
			return false
		}

		bigN = c
	}

	return true
}
