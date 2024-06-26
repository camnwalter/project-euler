package problems

import "github.com/camnwalter/project-euler/utils"

func ThirtyTwo() int {
	sum := 0

	// 1d x 4d = 4d
	// 2d x 3d = 4d
	pandigitals := make(map[int]bool)

	for x := 1; x <= 98; x++ {
		for y := 1; y <= 9876; y++ {
			product := x * y
			if !pandigitals[product] {

				digits := utils.ToDigitArray(x)
				digits = append(digits, utils.ToDigitArray(y)...)
				digits = append(digits, utils.ToDigitArray(x*y)...)

				if isPandigital(digits) {
					sum += product
					pandigitals[product] = true
				}
			}
		}
	}

	return sum
}

func isPandigital(digits []int) bool {

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
