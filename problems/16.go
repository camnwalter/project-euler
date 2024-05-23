package problems

import "github.com/camnwalter/project-euler/big"

func Sixteen() int {
	bigNumber := big.New(2)
	out := big.New(1)

	for range 1000 {
		out = out.Times(bigNumber)
	}

	sum := 0
	for _, digit := range out.Digits {
		sum += digit
	}

	return sum
}
