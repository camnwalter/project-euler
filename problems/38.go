package problems

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/camnwalter/project-euler/utils"
)

func ThirtyEight() int {
	max := 0
	for maxN := 2; maxN <= 9; maxN++ {
		for x := 1; x < 9876; x++ {
			var buf strings.Builder
			buf.WriteString(fmt.Sprint(x))

			for n := 2; n <= maxN; n++ {
				buf.WriteString(fmt.Sprint(x * maxN))
			}

			num, _ := strconv.Atoi(buf.String())
			if isPandigital(utils.ToDigitArray(num)) && num > max {
				max = num
			}
		}
	}

	return max
}
