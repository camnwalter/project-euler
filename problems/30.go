package problems

import (
	"math"

	"github.com/camnwalter/project-euler/utils"
)

func Thirty() int {
	n := 2
	sum := 0

	digits := 1
	maxDigitValue := utils.Pow(9, 5)

	// max = num digits * 9^5

	for n <= digits*maxDigitValue {
		partialSum := 0
		for _, digit := range utils.ToDigitArray(n) {
			partialSum += utils.Pow(digit, 5)
		}

		if partialSum == n {
			sum += n
		}

		n++
		digits = int(math.Log10(float64(n))) + 1
	}

	return sum
}
