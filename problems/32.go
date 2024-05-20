package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func ThirtyTwo() {
	sum := 0

	// 1d x 4d = 4d
	// 2d x 3d = 4d
	pandigitals := make(map[int]bool)

	for x := 1; x <= 98; x++ {
		for y := 1; y <= 9876; y++ {
			product := x * y
			if !pandigitals[product] && IsPandigital(x, y) {
				sum += product
				pandigitals[product] = true
			}
		}
	}

	fmt.Println("sum of pandigital products =", sum)
}

func IsPandigital(a int, b int) bool {
	digits := utils.ToDigitArray(a)
	digits = append(digits, utils.ToDigitArray(b)...)
	digits = append(digits, utils.ToDigitArray(a*b)...)

	if len(digits) != 9 {
		return false
	}

	seen := make([]bool, 9)

	for _, digit := range digits {
		if digit == 0 {
			return false
		}

		if seen[digit-1] {
			return false
		}
		seen[digit-1] = true
	}

	for _, v := range seen {
		if !v {
			return false
		}
	}

	return true
}
