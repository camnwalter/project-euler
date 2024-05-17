package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func Twenty() {
	fmt.Println("100! digits summed =", utils.SumArray(BigFactorial(100)))
}

func BigFactorial(n int) []int {
	out := utils.ToDigitArray(n)

	for i := 1; i < n; i++ {
		out = arrayTimes(out, i)
	}

	return out
}

// multiplies data * n
func arrayTimes(data []int, n int) []int {
	carry := 0

	for i := 0; i < len(data); i++ {
		product := data[i]*n + carry

		data[i] = product % 10
		carry = product / 10
	}

	for carry != 0 {
		data = append(data, carry%10)
		carry /= 10
	}

	return data
}
