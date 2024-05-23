package problems

import (
	"strconv"

	"github.com/camnwalter/project-euler/utils"
)

func ThirtySix() int {
	sum := 0

	for n := 1; n < 1_000_000; n++ {
		baseTen := strconv.FormatInt(int64(n), 10)
		baseTwo := strconv.FormatInt(int64(n), 2)

		if utils.IsPalindrome(baseTen) && utils.IsPalindrome(baseTwo) {
			sum += n
		}
	}

	return sum
}
