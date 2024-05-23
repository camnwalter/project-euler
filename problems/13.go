package problems

import (
	"slices"

	"github.com/camnwalter/project-euler/utils"
)

func Thirteen() int {
	lines, _ := utils.GetFileLines("inputs/13.txt")

	out := make([]int, 0)

	carry := 0
	for i := len(lines[0]) - 1; i >= 0; i-- {
		partialSum := carry
		for _, line := range lines {
			digit := int(line[i] - '0')
			partialSum += digit
		}

		out = append(out, partialSum%10)
		carry = partialSum / 10
	}

	for carry != 0 {
		out = append(out, carry%10)
		carry /= 10
	}

	for len(out) > 10 {
		out = slices.Delete(out, 0, 1)
	}

	sum := 0
	place := 1
	for _, digit := range out {
		sum += digit * place
		place *= 10
	}

	return sum
}
