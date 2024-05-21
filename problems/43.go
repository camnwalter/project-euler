package problems

import (
	"fmt"

	"github.com/camnwalter/project-euler/utils"
)

func FortyThree() {
	primes := PrimeSieve(18)

	sum := 0

	for _, pan := range GeneratePandigitals(0, 9) {
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

	fmt.Println("sum of pandigitals =", sum)
}

func GeneratePandigitals(min int, max int) []int {
	digits := make([]int, 0)
	for i := min; i <= max; i++ {
		digits = append(digits, i)
	}

	return utils.Map(utils.Permutations(digits), func(val []int, _ int) int {
		return utils.FromDigitArray(val)
	})
}
