package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func TwentyThree() {
	abundantNumbers := make(map[int]bool, 0)

	for i := 1; i < 28123; i++ {
		abundantNumbers[i] = isAbundant(i)
	}

	sum := 0

	for i := 1; i < 28123; i++ {
		abundantFound := false
		for j := 1; j <= i/2; j++ {
			if abundantNumbers[j] && abundantNumbers[i-j] {
				abundantFound = true
				break
			}
		}

		if !abundantFound {
			sum += i
		}
	}

	fmt.Println("sum of non-abundant nums =", sum)
}

func isAbundant(n int) bool {
	return utils.SumArray(GetProperDivisors(n)) > n
}
