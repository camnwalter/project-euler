package problems

import (
	"fmt"
	"strconv"
)

func ThirtyEight() {
	max := 0
	for maxN := 2; maxN <= 9; maxN++ {
		for x := 1; x < 9876; x++ {
			str := fmt.Sprint(x)

			for n := 2; n <= maxN; n++ {
				str += fmt.Sprint(x * maxN)
			}

			num, _ := strconv.Atoi(str)
			if IsPandigital(num) && num > max {
				max = num
			}
		}
	}

	fmt.Println("max pandigital =", max)
}
