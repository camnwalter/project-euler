package problems

import "github.com/camnwalter/project-euler/utils"

func ThirtyFour() int {
	// max = 9! * digits
	//     = 362880 * digits

	max := 7 * utils.Factorial(9)
	// 7*9! < 9_999_999

	sum := 0

	for n := 3; n < max; n++ {
		if digitFactorials(n) {
			sum += n
		}
	}

	return sum
}

func digitFactorials(n int) bool {
	digits := utils.ToDigitArray(n)

	sum := 0

	for _, digit := range digits {
		sum += utils.Factorial(digit)
	}

	return sum == n
}
