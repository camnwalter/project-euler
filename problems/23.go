package problems

import "github.com/camnwalter/project-euler/utils"

func TwentyThree() int {
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

	return sum
}

func isAbundant(n int) bool {
	return utils.SumArray(utils.GetProperDivisors(n)) > n
}
