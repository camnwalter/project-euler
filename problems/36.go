package problems

import (
	"fmt"
	"strconv"

	"github.com/camnwalter/project-euler/utils"
)

func ThirtySix() {
	sum := 0

	for n := 1; n < 1_000_000; n++ {
		baseTen := strconv.FormatInt(int64(n), 10)
		baseTwo := strconv.FormatInt(int64(n), 2)

		if IsPalindrome(baseTen) && IsPalindrome(baseTwo) {
			sum += n
		}
	}

	fmt.Println("sum of palindromic base 10 and 2 =", sum)
}

func IsPalindrome(num string) bool {
	return num == utils.Reverse(num)
}
