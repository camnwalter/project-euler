package problems

import (
	"github.com/camnwalter/project-euler/big"
	"github.com/camnwalter/project-euler/utils"
)

func SixtyFive() int {
	nums := make([]int, 99)
	k := 2
	for i := range 99 {
		num := 1
		if (i-1)%3 == 0 {
			// 2k
			num = k

			k += 2
		}
		nums[i] = num
	}

	a := big.New(0)
	b := big.New(1)
	for i := len(nums) - 1; i >= 0; i-- {
		b, a = fracAdd(big.New(nums[i]), a, b)
	}

	a, _ = fracAdd(big.New(2), a, b)

	return utils.SumArray(a.Digits)
}

// a + b/c
func fracAdd(a *big.MyBigInt, b *big.MyBigInt, c *big.MyBigInt) (*big.MyBigInt, *big.MyBigInt) {
	num, _ := a.Times(c).Plus(b)
	denom := c

	return num, denom
}
