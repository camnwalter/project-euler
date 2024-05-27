package problems

import "github.com/camnwalter/project-euler/big"

func FiftySeven() int {
	num := big.New(1)
	denom := big.New(1)

	count := 0

	for range 1000 {
		newDenom, _ := num.Plus(denom)
		newNum, _ := newDenom.Plus(denom)

		num = newNum
		denom = newDenom

		if len(num.Digits) > len(denom.Digits) {
			count++
		}
	}

	return count
}
