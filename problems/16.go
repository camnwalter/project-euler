package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/big"
)

func Sixteen() {
	bigNumber := big.New(2)
	out := big.New(1)

	for range 1000 {
		out = out.Times(bigNumber)
	}

	sum := 0
	for _, digit := range out.Digits {
		sum += digit
	}

	fmt.Println("sum =", sum)
}
