package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func ThirtyFour() {
	// max = 9! * digits
	//     = 362880 * digits

	max := 7 * Factorial(9)
	// 7*9! < 9_999_999

	sum := 0

	for n := 3; n < max; n++ {
		if digitFactorials(n) {
			sum += n
		}
	}

	fmt.Println("sum of digit factorial sums =", sum)

}

func digitFactorials(n int) bool {
	digits := utils.ToDigitArray(n)

	sum := 0

	for _, digit := range digits {
		sum += Factorial(digit)
	}

	return sum == n
}
