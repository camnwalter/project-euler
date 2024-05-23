package problems

import (
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func ThirtyThree() int {
	num := 1
	denom := 1

	for commonDigit := 1; commonDigit <= 9; commonDigit++ {
		for n := 1; n <= 9; n++ {
			if n == commonDigit {
				continue
			}

			n1 := n*10 + commonDigit
			n2 := commonDigit*10 + n

			for d := n + 1; d <= 9; d++ {
				if d == commonDigit {
					continue
				}

				d1 := d*10 + commonDigit
				d2 := commonDigit*10 + d

				if commonDigitReduction(n1, d1, commonDigit) {
					num *= n1
					denom *= d1
				}
				if commonDigitReduction(n1, d2, commonDigit) {
					num *= n1
					denom *= d2
				}
				if commonDigitReduction(n2, d1, commonDigit) {
					num *= n2
					denom *= d1
				}
				if commonDigitReduction(n2, d2, commonDigit) {
					num *= n2
					denom *= d2
				}
			}
		}
	}

	gcf := utils.GCF(num, denom)

	return denom / gcf
}

// a / b
func commonDigitReduction(a int, b int, common int) bool {
	if a >= b {
		return false
	}

	n := utils.ToDigitArray(a)
	i := slices.Index(n, common)
	n = slices.Delete(n, i, i+1)

	d := utils.ToDigitArray(b)
	i = slices.Index(d, common)
	d = slices.Delete(d, i, i+1)

	frac1 := float64(a) / float64(b)
	frac2 := float64(n[0]) / float64(d[0])

	return frac1 == frac2
}
