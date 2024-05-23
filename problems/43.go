package problems

import "github.com/camnwalter/project-euler/utils"

func FortyThree() int {
	primes := utils.PrimeSieve(18)

	sum := 0

	for _, pan := range generatePandigitals(0, 9) {
		index := 0
		good := true

		for i := 9; i >= 3; i-- {
			num := utils.GetDigits(pan, i-2, i)
			if num%primes[index] != 0 {
				good = false
				break
			}
			index++
		}

		if good {
			sum += pan
		}
	}

	return sum
}

func generatePandigitals(min int, max int) []int {
	digits := make([]int, 0)
	for i := min; i <= max; i++ {
		digits = append(digits, i)
	}

	return utils.Map(utils.Permutations(digits), func(val []int, _ int) int {
		return utils.FromDigitArray(val)
	})
}
